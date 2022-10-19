# 百度脑图研究（kityminder-core）

标签（空格分隔）： 百度脑图

---

## Kity Minder整体设计

### `KityMinder`

对应的 js：`kityminder.js`，实现了类似`namespace`的功能

暴露的命名空间，所有公开类都会放在该命名空间下。还会暴露一个简写的命名空间：KM

（项目中好像没找到，倒是它的例子里面将 Minder 的实例赋值给 KM，作为windows顶层变量来使用）
```
const Minder = KityMinder.Minder;
//创建km实例
const KM = window.KM = new Minder();
```

### `Command`

表示一条在 `KityMinder` 上执行的命令，以`class`的方式定义，**命令必须依附于模块，不允许单独存在**

我们先来看一下`Command`是怎么用的。比如我想在脑图上添加一个工作量，值为10，可以执行命令：
```
 KM.execCommand('efforts', 10);
```
当然，这个命令我们现在还没有去实现它。那么，如何去实现呢？

前面提到，命令必须依附于模块，不允许单独存在。所以在`KityMinder`里面，所有的命令都在`module`目录里面，包括它提供的或者我们自己定义的。
![image_1cn0mpu6d1f0e1ivqd981vqd5h39.png-23.1kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114063232.png)

所以在讲`Command`的时候我们来结合`Module`一起来讲

### `Module`

`Module`定义一个模块，表示控制脑图中一个功能的模块（布局、渲染、输入文字、图标叠加等）

#### 模块定义
那么我们如何去定义一个模块呢？在 `KityMinder`中有两种方式：

- 直接注册
    通过`Module`类的`register`方法直接注册一个模块，然后配置它的一些属性，实现它的一些方法
    ![image_1cn10bnp71sd06sh1g6713c4ohe16.png-44.9kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114064141.png)

- 通过`function(){return{}}`的形式
    对于一些比较复杂的模块，我们可以在`function`里面做一些操作，再去返回它的配置项。这样代码比较结构化一些，便于阅读。
    ![image_1cn10tnst1st2r1k1uaufkmb8e3g.png-50.3kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114065151.png)

#### 模块配置

在上面我们看到注册一个模块后，需要配置一些属性和实现一些方法，下面就来讲一下这些配置项是用来干嘛的。

它的官方文档里介绍了`defaultOpitons、init、commands、events、destroy和reset`这几个配置项，但在翻阅它`Module`里的文件时，还发现有`renderers、contextmenu、commandShortcutKeys`这些。

我们先看一下一个最基本的命令模块实现，比如我要执行一个下面的命令在节点上添加一个圆
```
//添加一个半径为2的圆
KM.execCommand('circle', 2);
```
模块的实现如下：（**参考`demo5`和`simple-command`**）
```
/**
 * 实现一个最简单的命令模块
 */
Module.register('addCircleModule', {
    'commands': {
        'circle': kity.createClass('Circle', {
            base: Command,
            execute: (minder, radius)=> {
                //获取选中的节点
                const node = minder.getSelectedNode();
                if (!node) {
                    return;
                }
                node.setData('radius', radius);
                node.render();
                minder.layout();
            },
            queryState: minder=>minder.getSelectedNodes().length === 1 ? 0 : -1,
            queryValue: (minder)=> {
                const node = minder.getSelectedNode();
                return node ? node.getData('radius') : null;
            }
        })
    },

    'renderers': {
        left: kity.createClass('CircleRender', {
            base: Renderer,
            create: function () {
                return new kity.Circle(0, 0, 0).fill('red');
            },
            shouldRender: function (node) {
                return node.getData('radius');
            },
            update: function (circle, node, box) {
                const radius = node.getData('radius');
                circle.setRadius(radius);
            }
        })
    }
});
```
这段代码实现了`commands`和`renderers`这两个方法

##### commands

`commands`就是我们前面说的命令了，所以通过代码我们就能理解命令必须依附于模块，不允许单独存在的意思了

`commands`里面可以包含多个命令，在上面代码中只有个`circle`命令。该命令的实现是创建一个`Command`类的子类。我们进入`Command`类，可以看到里面提供了一些方法，其中`execute`子类必须要实现。

所以我们在这里要自己实现`execute`方法，该方法的第一个参数为`minder`类的对象，后面的参数为我们执行命令时传入的参数，在这里我们传入了半径。

该方法就是命令的具体业务实现了，对于在节点上绘制一些东西，差不多都类似上面的写法。

然后还有两个缺省方法`queryState`和`queryValue`

`queryState`用于返回当前命令执行的状态

- -1：不可执行
- 0：可执行
- 1：已执行

`queryValue`用于返回当前命令的状态相关值，（例如：进度条的进度百分比值等）

##### renderers

`renderers` 部分需要实现该命令要在节点上绘制的内容了。

它里面的键`left`、`right`不知道干吗用的，目前还没发现用处

它的值是一个`Renderer`类的子类。同样，我们进去`Renderer`类，可以发现有`create`、`draw`和`place`三个方法子类要自己实现

`draw`和`place`方法，我们可以用`update`方法来代替，因为它内部已经实现了这两个方法

然后，大部分情况我们还要实现一个缺省方法`shouldRender`来判断是否要绘制

`create`方法就是图形的具体绘制了，通常我们会返回一个图形对象。然后在`update`方法里做一些更新操作，比如：拿到命令传递的参数为图形设置相应的属性。通常`update`方法最后还会返回一个新的`Box`，这个后面再做介绍。

大部分情况下，上面这两个配置项就可以完成了。

##### defaultOpitons

`defaultOpitons`可以配置模块可能使用到的一些默认值。比如在它`image.js`里面，默认配置图片的最大宽高：
```
'defaultOptions': {
    'maxImageWidth': 200,
    'maxImageHeight': 200
},
```
然后在`execute`方法中通过`minder.getOption('maxImageWidth')`和`minder.getOption('maxImageHeight')`就可以取到值，做一些操作

##### init
`init`官方的解释是：
```
// Minder 实例化的时候会调用 init 方法，this 指向正在实例化的 Minder 对象
// options 是 Minder 对象最终的配置（经过配置文件和用户设定改写）
"init": function( options ){
},
```
这个方法里面的内容脑图实例化的时候就会去运行的，比如在`image-viewer.js`中
```
init: function () {
    this.viewer = new ImageViewer();
},
```
在`init`里实例化一个`ImageViewer`对象，然后在`events`对应的事件里面去执行`this.viewer.open(shape.url);`方法就可以完成双击图片查看的功能了。

>但是这个最好不要多用，会影响脑图的加载。

##### destroy 和 reset
这两个方法系统中没有地方用到，暂时不清楚
```
// Minder 被卸载的时候会调用 destroy 方法，模块自行回收自己的资源（事件由 Minder 自动回收）
// destroy 方法中的 this 指向 Minder 实例
//By jerry.mei：该方法系统中没有使用到，暂时不知道用处
"destroy": function () {
},
// Minder 被重设是会调用 reset 方法，模块自行
// reset 方法中的 this 指向 Minder 实例
// by jerry.mei：该方法系统中没有使用到，暂时不知道用处
"reset": function () {
}
```

##### events

事件的绑定，可以为添加的图形绑定事件

```
'events': {
    'click': function (e) {
        // console.log(e);
    },
    'keydown keyup': function (e) {
        // console.log(e);
    }
},
```

##### commandShortcutKeys

快捷键功能，脑图里绑定快捷键十分方便，只需要将命令与对应的快捷键映射一下，就支持快捷键操作了
```
'commandShortcutKeys': {
    'efforts': 'ctrl+alt+c',
}
```

##### contextmenu
暂时不清楚

### `MinderNode`

### `Minder`
脑图使用类

通过初始化 `Minder` 类，在`constructor`里面会创建脑图画布
```
const KM = window.KM = new Minder();
```
 
将 `KM` 暴露为浏览器顶层对象，后续的大部分操作都可以使用 `KM` 来完成


## Box 在 KityMinder 里的使用
>参考`demo5`和`simple-command2`

### Box详解
之前讲kity的时候着重讲了一下`Box`，因为在`KityMinder`里面，`Box`决定了节点里面元素的排列位置

为了更方便地了解`Box`在`KityMinder`里的使用，我写了一个自定义的主题，将默认的padding什么的都设置为0

我们切换到自定义主题，添加一个节点，发现就是这样

![image_1cn1cd73r1b011269cgu1ok51qt4h.png-47.4kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114065959.png)

获取它的`Box`属性

![image_1cn1cg550t563oj4lfau61u1n5e.png-13.3kB](https://image.maplejoyous.cn/post/2022/10/11/202210111407066.png)

x,y为左上角顶点坐标（0，-7），所以我们可以得出该节点坐标系原点的位置。放大看如下图：
所以其他的一些属性也很好解释

![image_1cn1ctvqh4ri1h07s6818ib18te6u.png-3.3kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114071414.png)

所以当我们往节点添加一个半径为5，坐标为（0,0）的圆时，这个圆在这个节点中是居中的吗？

![image_1cn1eqk1ivlg17j51k1v9ojdbl84.png-7.5kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114072222.png)

我们将浏览器放大到到500%，可以看到圆心在坐标轴的原点，并不在节点的中心点

同样，我们添加一个正方形，发现正方形的左上角顶点在坐标轴的原点。
![image_1cn1f43uqdhphlg1l7o1b3f1slm9k.png-5.8kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114072828.png)

所以，我们确定好坐标系后，就很容易去确定节点上元素的位置了，你只需要关心你绘制图形的顶点坐标位置就可以了。

比如圆形，我们看它的参数解释：

![image_1cn36hb8v66fhce1femrs4dnh1c.png-19.9kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114073636.png)

所以，我们设置它的坐标（0,0），其实是圆心在坐标系（0,0）的位置

再来看矩形：

![image_1cn36mqci1u2aldp1buc1btaren3o.png-44kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114074545.png)

我们设置它的坐标（0,0），其实是左上角顶点坐标在坐标系（0,0）的位置

### 利用 Box 确定元素位置

#### update 方法参数说明

当我们在脑图上新增一个节点后，默认会插入一个`Text`，它会有默认的`padding`等样式属性。

这些属性值是在`theme`目录里面配置的，比如我自定义了一个`custom.js`，里面会针对不同的节点去设置样式。

然后我们回到之前讲的`Module`里面`renderers`需要在`update`方法里返回一个新的盒子

![image_1cn39cjdh13771ekc1nao19cg1g6418.png-31.4kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114075353.png)

我们可以看到`update`方法里有三个参数，第一个参数为`create`方法返回的图形，第二个参数为选择的`node`节点，第三个参数为`box`

我们可利用`node`的`getStyle()`方法获取一些样式值，比如`space`属性，可以获取每个元素之间的间距，`padding`可以获取节点的内边距

我们打印一个`box`的值，因为节点默认添加了`Text`，所以`box`返回的是该`Text`的`box`值

#### 设置图形的位置

在`update`方法里面，我们可以用两种或者更多的方式来设置图形的位置：

- 改变图形的坐标
- 利用`kityminder`坐标系的`setTranslate`方法

这些都不是重要的，重要的是我们要获取我们需要摆放图形的位置

比如：我要在文字的右边添加一个圆形，如下图：

![image_1cn3b5ep61amf12dv1nrt18g8kqp35.png-26.9kB](https://image.maplejoyous.cn/post/2022/10/11/202210111408000.png)

我们可以获取文字 Box 的`right`值，即文字的最右边`x`轴坐标，加上`space-right`值即可
```
const x = box.right + node.getStyle('space-right');
//利用改变坐标设置图形的位置
circle.setCenterX(x);
//利用setTranslate方法设置图形的位置
circle.setTranslate(x, 0);
```

#### 返回新的盒子

在上面的操作之后，我们发现节点中新增了一个圆，并在文字的右边，好像符合了我们的预期，但其实并不是这样。

因为我们给 `node` 节点添加了默认的`padding`属性，看起来好像文字和圆都在节点里面。但当我们去掉 `node`默认的`padding`值，问题就出现了。

我们设置一下自定义主题，看到的结果如下：

![image_1cn3bts3tlne1dqf1f29lhp75p3i.png-23.1kB](https://image.maplejoyous.cn/post/2022/10/11/202210111408099.png)

因为我们的盒子还是之前 text 的盒子，添加了圆之后我们需要返回新的盒子：

```
//返回新的盒子
return new kity.Box({
    x: x,
    y: box.y,
    width: circle.rx,
    height: circle.ry
})
```
x,y 即圆心的坐标，宽高为圆的 x,y 半径

我们去掉样式后看到的效果就是这样的：

![image_1cn3ilfp31nbi1fq21fdlv6hgjc3v.png-20.8kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114081515.png)

## kityminder-core 文件详解

我们可以看到`kityminder-core`跟`kity`的文件组织方式类似。

`kityminder.js`用于导出所有的类和模块，`expose-kityminder.js`用于将`kityminder`变量提升为`windows`对象

![image_1cn5ukkrp1egtep4t5behdrbf9.png-10.4kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114082323.png)

因为我们需要去扩展`kityminder-core`的类库，所以将`kityminder-core`作为项目的一个模块，去掉了`expose-kityminder.js`，使用的时候直接引入`kityminder.js`即可

### core部分

**`utils`**

工具文件，提供了一些常用的方法

**`minder`**

比较重要的一个文件，使用了 initHooks[] 机制, 对每个注册的 init-hook 用 this 产生调用.

似乎还有 event 机制, 在 ctor 最后使用 fire() 触发 'finishInitHook' 事件.

脑图的大部分文件都是通过拓展`Minder`类，来完成相应的业务的。

注释里面说，是暴露在 window 上的唯一变量

**`command`**

通过 `execCommand` 方法来执行一些命令的操作
```
km.execCommand('HyperLink', 'https://www.baidu.com', '百度');
```

按照架构文档说明, abstract Command 表示一条在 KityMinder 上执行的命令.
以class的方式定义，命令必须依附于模块，不允许单独存在。

这个文件做了两件事情：一个是实现了一个`Command`类，还有就是拓展了`Minder`类。接下来的大部分文件核心内容都是在做这两件事。

先看一下 `Command` 类实现的部分：

`execute` 方法，执行一条命令。类似于一个抽象方法，子类必须实现

然后就是一些 set/get 方法了

再来看一下拓展`Minder`类的部分，实现了如下的方法：

- _getCommand
    获取命令的相关信息，输出一下`KM._getCommand('HyperLink')`
    ![QQ截图20180903151541.png-11.3kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114083030.png)
- _queryCommand
- queryCommandState
    获取当前命令的装填
- queryCommandValue
- execCommand

**`node`**

创建脑图的一个节点

**`event`**

表示一个脑图中发生的事件

这里也回答了前面的问题, 即 Minder 的 event 从哪里来的问题.

## 脑图加载的执行顺序

在脑图的 demo 中，我们先是在 HTML 中写了一段 script，内容就是脑图的 json 对象
```
<script id="minder-view" type="application/kityminder" minder-data-type="json">
    {
        "root": {
            "data": {
                "text": "百度产品"
            },
            "children": [
                { "data": { "text": "新闻" },"children":[
                    {"data":{"text":"财经板块"}},
                    {"data":{"text":"文娱板块"}}
                ] },
                { "data": { "text": "地方都是地王大厦\n第三方的\n第三方" }},
                { "data": { "text": "贴吧", "priority": 2 } },
                { "data": { "text": "知道", "priority": 2 } },
                { "data": { "text": "音乐", "priority": 3 } },
                { "data": { "text": "图片", "priority": 3 } },
                { "data": { "text": "视频", "priority": 3 } },
                { "data": { "text": "地图", "priority": 3 } },
                { "data": { "text": "百科" }},
                { "data": { "text": "百度脑涂" } },
                { "data": { "text": "更多", "hyperlink": "http://www.baidu.com/more" } }
            ]
        }
    }
</script>
```
然后执行`KM.setup('#minder-view')` 就可以将脑图加载出来了。那么，脑图的整个加载流程是怎样的呢？

首先，`setup` 方法通过 HTML DOM 操作，根据传进来的参数获取 DOM 对象以及 DOM 对象的属性和内容。

获取到了`protocol`后，判断是否在`protocols`中，如果在就依次执行`renderTo`和`importData`方法，渲染节点，导入数据。

我们在调试的时候，可以发现`protocols`对象的值如下：
![image_1cn3p75i21s3b1mph7d01dod11199.png-35.8kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114084040.png)

它默认支持了这些格式的数据。那么`protocols`是在哪里赋的值呢？

我们可以发现在`data.js`中一开始就定义了`protocols`对象，然后实现并导出了两个方法：
```
var protocols = {};

function registerProtocol(name, protocol) {
    protocols[name] = protocol;

    for (var pname in protocols) {
        if (protocols.hasOwnProperty(pname)) {
            protocols[pname] = protocols[pname];
            protocols[pname].name = pname;
        }
    }
}

function getRegisterProtocol(name) {
    return name === undefined ? protocols : (protocols[name] || null);
}

exports.registerProtocol = registerProtocol;
exports.getRegisterProtocol = getRegisterProtocol;
```

`registerProtocol`方法是关键，它其实就是在给`protocols`添加对象。那么，我们只需要知道哪里调用了`registerProtocol`方法，就可以知道`protocols`的值是怎么产生的了。

搜索发现，在`protocol`目录下的这些文件都是在给`protocols`添加对象的。

![image_1cn3psgjpa7h10r57l911m4125626.png-5.7kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114084747.png)

我们也可以参照它的实现去扩展自己的一些数据格式。

因为我们在 `kityminder.js`中引入了这些模块依赖，所以这些方法一开始就都会执行的。`protocols`的内容就是这样来的。

接下来，我们看`renderTo`方法的实现。

`renderTo`方法是在`paper.js`里的，在`kity`里我们知道`paper`是画布，一切的显示最终都是要在画布上呈现。

这个文件一开始使用了`initHook`的机制去执行了`_initPaper`方法，做一些 paper 初始化的操作
```
Minder.registerInitHook(function () {
    this._initPaper();
});
```
在`renderTo`方法的内部，其实就是利用 kity paper 的 `renderTo`方法将 paper 渲染到真实的 DOM 节点上了。；然后`_bindEvents`方法为 paper 绑定了一些事件；最后`this.fire('paperrender')`触发了一个`paperrender`事件。

关于`fire`，前面 `Minder`里有介绍到。它仅仅是触发了一个事件，就像你点击一个按钮触发了click事件，但事件的回调我们需要在具体的地方去实现。比如：在它的`view.js`里

![image_1cn3svicm4lb2ehongr21p9589.png-17.3kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114095151.png)

到这一步，DOM 中已经插入了 id为`minder-view`的 div，里面 svg 渲染了一个根节点，在画布的最左侧

![image_1cn3tauodnqu1iopur81bkf10uje1.png-19.7kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114095959.png)

改一下它的位移，发现它是这样的：

![image_1cn3tf5jml8i1941kq3hgcu0gih.png-1.4kB](https://image.maplejoyous.cn/post/2022/10/11/2022101114125959.png)

`importData`方法是根据传入的脑图数据来开始绘制脑图了。我么看`importData`方法的实现。

这里面首先是对协议进行检查，然后这个导入前抛事件`this._fire(new MinderEvent('beforeimport', params))`不知道干吗的，系统中也没用到，文档也没提到。

`return` 这部分比较重要，通过 js 的 `Promise`来实现一个异步操作，异步操作里面执行`importJson`方法，并返回 `json` 数据

所以最关键的部分就是`importJson`了，这里面完成绘制的操作。

我们看`importJson`方法：

一开始也是使用`_fire`触发一个前置事件，然后删除当前的所有节点。

`importNode`方法是绘制的核心，将 json 对象绘制成对应的脑图节点

这里面主要是循环去做`setData`和`createNode`，`setData`将节点的数据传入进去，`createNode`就是创建节点了。

然后节点上图形的绘制就是`module`里的那些模块利用`renderers`实现的。

节点的连线是`connect`里提供的一些连线函数来实现的

这些都做完了后，通过下面的代码来设置主题和布局
```
this.setTemplate(json.template || 'default');
this.setTheme(json.theme || null);
this.refresh();
```


  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  
  