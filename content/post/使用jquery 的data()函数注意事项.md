+++
date = "2015-11-04T17:12:34+08:00"
title = "使用jquery 的data()函数注意事项"
tags = [ "jquery", "data" ]
categories = ["前端"]
+++

<div data-kindType="1" id="test"></div>

这种写法在jquery缓存中保存的key并不是kindType，而是kind-type，故此$('#test').data('kindType')是取不到值的。此处要注意