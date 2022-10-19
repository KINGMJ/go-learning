# 百度脑图研究（kity）

标签（空格分隔）： 百度脑图

---

## 百度脑图总体架构

![image_1cmh88b21po1momrm3r3tc4t9.png-437.5kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114161313.png)

- [百度脑图](http://naotu.baidu.com/) 基于 kityminder-editor，加入了第三方格式导入导出 (FreeMind, XMind, MindManager) 、文件储存、用户认证、文件分享、历史版本等业务逻辑。属于一个比较完整的产品。

- [kityminder-editor](https://github.com/fex-team/kityminder-editor) 基于 kityminder-core 搭建，依赖于 AngularJS，包含 UI 和热盒 hotbox 等方便用户输入的功能，简单来说，就是一款编辑器。

- [kityminder-core](https://github.com/fex-team/kityminder-core) 是 kityminder 的核心部分，基于百度 FEX 开发的矢量图形库 kity。包含了脑图数据的可视化展现，简单编辑功能等所有底层支持。

- [kity](https://github.com/fex-team/kity) 是一个基于 SVG 的矢量图形库，帮助你快速在页面上创建和使用矢量元素。类似于 D3 这种库，偏向于底层。


## Kity

### 基本介绍

Kity 是一款矢量图形库，它的最底层特性是面向对象支持。众所周知JavaScript是弱类型语言，要想用它完成工程化的产出，就需要面向对象的支持。因此我们基于ES5构建了一套面向对象语言的实现，包括定义类、扩展类、继承以及混合。下图就是Kity类的继承关系。

![image_1cmh943uffn2csqmbo15hf1g9s26.png-202.2kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114162323.png)
**`类图参考价值不大，比较老`**

**Kity的核心如下：**

- 图形库
- 坐标系
- OOP

### 项目文件解析

![image_1cmha9b9a1crc1envdugtc117d2j.png-8.7kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114163030.png)

kity 的整个文件结构如图所示，包括：

- `kity.js` 用于导出所有模块
- `expose-kity.js` 将`kity`变量提升为 `windows`对象
- `graphic` 图形库，实现各种点、线、面
- `filter` 滤镜库，实现一些高斯模糊等效果
- `core` 核心库
- `animate` 动画相关

`kity.js` 主要使用 `seajs` 的模块系统将整个项目需要使用到的方法或对象暴露出来，然后用`kity`来作为一个顶级对象，实现了类似于`namespace`的功能。

`expose-kity.js` 的作用如上所示，对于普通引入方式，引入了生产环境的`kity.js`文件就可以直接使用`kity`变量了；对于用npm方式安装的直接`import 'kity'`即可。

`core` 里面包含 `browser.js`、`class.js`和`utils.js`。

`browser.js` 主要提供了一些浏览器判断的方法，通过输出`kity.Browser`我们可以看到当前浏览器的一些属性。
![image_1cmhh8vm81h10td1oj2v62b0i9.png-19.9kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114163737.png)

`utils.js` 见名知意，提供了一些常用的方法，如下：
![image_1cmhhbrqopoquti1rmo6a51uqrm.png-53.1kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114164545.png)

`class.js` 这个比较重要，提供了 kity 的 oop 支持，class的用法可以参考 `demo4`，
另外 `pipe`函数的使用可以参考`demo1`和`demo2`

### graphic 图形系统

在 `graphic` 文件夹里面，有很多图形相关的文件。根据 [部分文档](https://github.com/fex-team/kity/wiki) 提供的信息，首先是画布与图形。

#### 画布（参考`demo6`）

在 Kity 中，画布用 Paper 表示，它是最基本的绘图容器。我们进入到`paper.js`中，看一下它的实现：

在`constructor` 函数中，使用 `createSVGNode` 函数创建了 svg 的节点，然后使用 `renderTo` 方法渲染到指定的 HTML DOM 元素上

##### 初始化 Paper

用下面的代码可以创建一个 Paper，并且会在指定的容器中渲染：
```
const paper = new kity.Paper('container')
// 或者直接传 Dom 对象：
const paper = new Paper(document.body);

```

执行可以看到 container 下已经有一个 svg 标签了
![image_1cmhj2e2818mlrb64l01nmo72s1j.png-10.5kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114165353.png)

##### 设置宽高和视野（参考`demo6`和`demo12`）

在 paper 中，有很多方法，比如设置宽高：
```
paper.setWidth(800)
    .setHeight(600);
```
设置视野：

- **`setViewPort`**

设置视窗。paper自己的方法，参数说明：（x坐标，y坐标，transform的放大缩小系数）
x=0,y=0时，视窗在paper的中心点处，矩形看起来只有实际大小的一半，如图：
```
paper.setViewPort(0, 0, 0.5);
```
![image_1cmkgua07jv51u6vhhpjb31b299i.png-12.9kB](https://image.maplejoyous.cn/post/2022/10/11/202210111417000.png)

通过`getViewPort`方法可以获取如下信息：
![image_1cmkh7gbc1juv1aaf1mcf2jksdj9v.png-7.9kB](https://image.maplejoyous.cn/post/2022/10/11/202210111417088.png)

`center`值都为0；
`offset`值为400和300，因为paper的宽高为800和600，所以offset值对应了视窗偏移点的坐标
`zoom`值为0.5，缩小一半

所以前面两个参数有点像在确定Paper原点的坐标。



- **`setViewBox`**

属于`ViewBox`类的方法，跟`setViewPort`方法类似但有很大差别

```
paper.setViewBox(-800, -600, 1600, 1200);
```

视野定义了 Paper 下图形的坐标系统。由四个值来定义：( x, y, width, height )。其中 x 和 y 确定 Paper 左上角的点在坐标系里的坐标，而 width 和 height 就表示 Paper 显示的坐标范围。

`ViewBox`的理解感觉跟`ViewPort`是反着来的。在`ViewPort`中设置（0,0,0.5)，我们可以理解为将图形移动到了画布的中心点，并缩小了一半。 

但用`ViewBox`完成这种效果，我们可以理解为图不动，画布在动。需要缩小图形为一半，我们就要将画布
放大一倍；同样的，需要设置图形的顶点为画布中心点，我们需要将画布往负方向移动相应的距离。

**但实际上画布和图形都没有任何变化**

在 `paper.js` 中，我们并没有看到`setViewBox` 方法，这里是用了 `mixins` 来实现的。
```
 mixins: [ShapeContainer, EventHandler, Styled, ViewBox],
```
`Paper`类混合了上面的这些类，比如上面提到的`setViewBox`方法是在`viewbox.js`中实现的。

通过它们我们可以完成放大、缩小，图形在画布中位置的摆放等操作。

##### 图形管理（参考`demo1`）

因为 `Paper` 类混合了`ShapeContainer`类，所以它还是一个图形容器，可以向其添加和移除图形

我们可以看到`ShapeContainer`类中实现了很多关于图形操作的方法，通过调用这些方法可以在paper上任意添加、移除图形.

##### 事件

继续看 `paper.js`，里面还混入了`EventHandler`类，来实现paper的一些事件。
```
paper.on('click', function (e) {
    const mouse = e.originEvent;
    //距离窗口的位置偏移
    console.log(mouse.clientX);
    console.log(mouse.clientY);
    //距离画布的位置偏移
    console.log(mouse.offsetX);
    console.log(mouse.offsetY);
});
```

##### 样式管理
`paper.js`里面混入的`Styled`类是实现一些样式操作的。这个类比较简单，就实现了如下四个方法：

- `addClass`
- `removeClass`
- `hasClass`
- `setStyle`

##### 资源管理

在`paper.js`里面，`Paper`类自己实现了一个`addResource`方法，可以向画布上添加一些需要使用的资源。比如：`LinearGradientBrush`、`RadialGradientBrush`、`PatternBrush `等。

往 Paper 添加和移除资源使用以下接口：
```
//资源操作
const brush = new kity.LinearGradient().pipe(function () {
    this.addStop(0, new kity.Color('red'));
    this.addStop(1, new kity.Color('blue'));
});
paper.addResource(brush);
rect.fill(brush);

// 资源被移除后，矩形的填充会失效
paper.removeResource(brush);
```

#### 图形
Kity Graphic 内置了 `Path`、`Rect`、`Ellipse`、`Circle`、`Polyline`、`Polygon`、`Curve`、`Bezier` 等基本几何图形。

`Shape`类是所有图形的基类，里面提供了一些获取、设置属性，图形变换等方法。同样，它也混入了一些类来完成一些其他的操作：
```
mixins: [EventHandler, Styled, Data]
```

##### Path
在 svg 中，`<path>`元素用于定义一个路径。通过路径可以绘制出任何图形。

所以`Path`类 是 Kity 中最强大的工具。其他的几何图形都是继承 Path 而来。Path 能识别 SVG 中定义的 [Path Data](https://www.w3.org/TR/SVG/paths.html#PathData) 字符串格式。可以通过这样一个字符串构造 Path：
```
//绘制一个封闭的红色三角形
const triangle = new kity.Path('M150 0 L75 200 L225 200 Z').fill('red');
```

##### Group（参考`demo8`）
使用 group 可以建立图形分组，将一些图形组合。

然后可以直接对图形组进行图形操作。因为组本身也是一个图形（由其子元素组合），所以也可以被添加到组里。

##### 填充图形（参考`demo9`）

默认添加到 Paper 上的图形是不具有视觉呈现的，需要对其进行填充或描边。

- 纯色填充，使用 Color

    要用一个颜色进行填充，可以：
    ```
    rect.fill( new Color('red') );
    // 或者直接使用字符串：
    rect.fill( 'red' );
    
    ```
- 各种笔刷填充，利用kity内置的各种渐变和图形笔刷填充图形

    - 使用 `LinearGradientBrush` 进行线性渐变填充
    
    ```
    rect.fill(new kity.LinearGradient().pipe(function () {
    //添加关键颜色到具体位置，其中 0 表示渐变开始的位置，1 表示渐变结束的位置
    //参数说明：（位置，颜色，透明度）
    this.addStop(0, 'red', 1);
    this.addStop(0.5, 'yellow', 0.5);
    this.addStop(1, 'blue', 1);
    //设置填充的方向和范围，参数为坐标系坐标
    this.setStartPosition(0, 0);
    this.setEndPosition(0, 1);
    paper.addResource(this);
    }));
    ```
    ![image_1cmjp6l9itv87j818dd1jlj1rr316.png-2.7kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114171515.png)

    - 使用 `RadialGradientBrush` 进行径向渐变填充
    
    ```
    rect.fill(new kity.RadialGradient().pipe(function () {
    this.setCenter(0.5, 0.5);
    this.setRadius(0.5);
    this.setFocal(0.8, 0.2);
    this.addStop(0, 'red');
    this.addStop(1, 'blue');
    paper.addResource(this);
    }));
    ```
    ![image_1cmjp8v6d1nja1v2obe1tlgpji1j.png-26.3kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114172424.png)
    
    - 使用 `PatternBrush` 进行图形填充，`PatternBrush` 是最灵活的画笔，它可以用图形填充图形。
    
    ```
    rect.fill(new kity.Pattern().pipe(function () {
    const colors = ['red', 'blue', 'yelow', 'green'];
    this.setWidth(40);
    this.setHeight(40);
    this.addItem(new kity.Circle(5, 10, 10).fill(colors.shift()));
    this.addItem(new kity.Circle(5, 30, 10).fill(colors.shift()));
    this.addItem(new kity.Circle(5, 10, 30).fill(colors.shift()));
    this.addItem(new kity.Circle(5, 30, 30).fill(colors.shift()));
    paper.addResource(this);
    }));
    ```
    ![image_1cmjpcdc814j4uv2br2ihhl0n3g.png-6.2kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114173232.png)
    
    我们去看`PatternBrush`的源码，可以发现它混入了`ShapeContainer`类，所以它也是一个图形容器，我们可以向它里面添加一些图形。
    
    - 使用Pen 来绘制图形轮廓
    
        实现了 svg 的 [stroke 属性](http://www.runoob.com/svg/svg-stroke.html) 
    
    ```
    rect.stroke(new kity.Pen().pipe(function () {
    this.setWidth(1); //设置画笔的粗细
    this.setDashArray([5, 5]); //设置画笔的段长和间隙长，不断循环。默认为 null，绘制实线
    this.setLineCap('butt');   //设置端点的样式，取值有：butt、round、suqare
    this.setLineJoin('round');  //设置转折点的样式，取值有：miter、round、bevel
    this.setColor('green');
    }));
    ```
    ![image_1cmjr5p4o1cfcnvf1gt4iaa1rsu3t.png-2.6kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114174242.png)

##### 文字（参考`demo10`）

在 kity 中，与文字相关的有`text.js`、`textcontent.js`和`textspan.js`这三个文件。

我们可以在`kity.js`中看到默认暴露了`Text`类和`TextSpan`类，`TextContent`类作为这两个类的基类，实现了文字操作的一些方法。然后`TextContent`类的基类又是`Shape`，所以他们都可以当做图形来处理。

先看`TextSpan`类，它的实现很简单，就是在 svg 中创建了一个 [tspan](https://developer.mozilla.org/zh-CN/docs/Web/SVG/Element/tspan) 标签，然后设置内容。

既然它们都是图形类的子类，我们可以将它添加到画布上
```
const textSpan = new kity.TextSpan('hello world').fill('red');
paper.addShape(textSpan);
```

画布上没有出现任何东西，但是在 HTML 中，我们可以看到已经渲染出了 `tspan`标签
![image_1cmjt5f8osa11cgm1a4rj06v7m59.png-7.2kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114175151.png)

查看 MDN 关于 tspan 的解释，我们得知它是内含在 text 元素中，所以直接使用是没有任何效果的。

所以我们来看一下`Text`类，它的构造函数跟 TextSpan 差不多，我们可以将它添加到 Paper 上看一下效果

```
const text = new kity.Text('hello world').pipe(function () {
    this.setX(100);
    this.setY(100);
    this.setSize(36);
    this.fill('red');
});
paper.addShape(text);
```

画布上显示了 hello world，html元素里出现了`text`标签
![image_1cmju17s3371c4e1d4f1f6cnu27e.png-4.3kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114175959.png)
![image_1cmju32lm1k2comqjnqpkuocu7r.png-8.7kB](https://image.maplejoyous.cn/post/2022/10/11/202210111418077.png)

那么，`TextSpan` 与`Text`的关系是怎样的呢？

前面我们知道 `tspan` 标签是在 `text` 标签里面。我们去看`Text`类的实现，发现它混入了`ShapeContainer`类，所以它可以作为一个图形容器，添加`TextSpan`进去。它提供了一个`addSpan`方法，就是来完成这件事的。

```
const text = new kity.Text().pipe(function () {
    this.addSpan(new kity.TextSpan('hello').fill('red'));
    this.addSpan(new kity.TextSpan('world').fill('blue'));
    this.setX(100);
    this.setY(100);
    this.setSize(36);
});
```
![image_1cmk4f66q14kmj2j19ft1n12llj8o.png-4.7kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114181616.png)
![image_1cmk4g7fvh5j1abn1d3m1vai1faf95.png-15.9kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114182323.png)

另外，使用`setPath`方法可以使用文本路径来排列文字

### 内置图形介绍（参考`demo7`）

前面我们接触到了 Rect 矩形和使用 Path 来绘制自定义图形，在 kity 中还内置了很多丰富的图形。

#### `Ellipse`
`Ellipse` 用于绘制一个椭圆，参数解释：（x轴半径，y轴半径，x轴坐标，y轴坐标）
如果将x轴半径和y轴半径设置一样，就是一个圆形了
```
const ellipse = new kity.Ellipse(100, 60, 400, 150).fill('green');
paper.addShape(ellipse);
```

#### `Circle`
`Circle` 用于绘制一个圆形，参数解释：（半径，x轴坐标，y轴坐标）

#### `HyperLink`
`HyperLink` 用于生成一个链接，也就是 a 标签。但 svg 中的 a 标签与 HTML 中还是有很大区别的，如果直接像下面这种写法，是无效的。
```
<svg>
    <a href="http://www.baidu.com">百度</a>
</svg>
```

所以我们可以看到`HyperLink`类是混入了`ShapeContainer`类的，也就是说它是一个图形容器。我们可以在它里面添加一些图形，让该图形变成一个链接，比如`Text`：
```
const hyperlink = new kity.HyperLink('https://www.baidu.com').pipe(function () {
    this.addShape(new kity.Text('百度').pipe(function () {
        this.setX(300);
        this.setY(300);
        this.setSize(36);
        this.fill('blue');
    }));
    this.setTarget('_blank');
});
paper.addShape(hyperlink);
```

#### `Image`
图片的使用十分简单，初始化参数为：（url，宽，高，x轴坐标，y轴坐标）
```
//Image 图片
const image = new kity.Image('http://img2.ph.126.net/vZ1r-IVk_j-cAi-noOuCEw==/760826862149437991.png', 200, 200, 100, 400);
paper.addShape(image);
```

### 其他重要属性和方法介绍（参考`demo11`）

#### `getBoundaryBox`
`Shape`基类的方法，获取图形的边界盒子，返回一个 `Box`对象

#### `setBox`
`Rect`类的方法，设置一个盒子

#### `expand`
`Box`类的方法，扩展（或收缩）当前的盒子，返回新的盒子

#### `Box`属性介绍

我们添加一个`Text`，设置坐标为（100,100），并使用它的边界盒子属性绘制一个矩形。
控制台打印一下Box的属性：
![image_1cn0iu1rt15ut1mla1oa01j0l1srb2f.png-16.1kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114182929.png)

我们用鼠标去选择文字，可以看到选中的区域刚好就是该矩形的区域。

下面主要介绍这些属性值：

首先，我们知道`Text`的坐标为（100,100）,但我们可以看到x的值为100，y值确是83。

那这个83是怎么来的呢，我们用尺子量一下就知道了。

![image_1cn0j889b1lro1nbu1mra1hn9qjq5i.png-18.6kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114183737.png)

我们先纵向量100的距离，可以发现（100，100）刚好是文字底部的点坐标。

所以，我们不难发现（100,83）为该`Text`盒子的顶点坐标。

宽高就不用解释了，即盒子的宽高

四个点坐标：`top`、`bottom`、`left`、`right`

`top:83`，`left:100`即盒子的顶点坐标，（100,83）

`right= left+width`
`bottom=top+height`

`cx`、`cy`为矩形区域的中心点坐标，所以：

`cx=left+width/2`
`cy=top+height/2`


  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  