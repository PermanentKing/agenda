# agnda 使用

------

该包在cobra的基础上实现，请先安装好cobra包之后再进行使用。
下载该项目到go工作目录的src里面，然后运行`go install ageda`，即可生成agenda可运行程序。
使用过程如下：

1、 查看各种命令

![此处输入图片的描述][1]


2、 注册用户

![此处输入图片的描述][2]


3、 用户登录

![此处输入图片的描述][3]


4、 创建会议

![此处输入图片的描述][4]


5、 添加成员

![此处输入图片的描述][5]


6、 查询会议

![此处输入图片的描述][6]


7、 删除成员

![此处输入图片的描述][7]


8、 删除会议:删除之后再查询已经没有会议

![此处输入图片的描述][8]


9、 清空会议:清空之后再查询已经没有会议

![此处输入图片的描述][9]


10、 查询用户

![此处输入图片的描述][10]


11、 logout:登出之后不能进行相关操作

![此处输入图片的描述][11]


12、 删除账户：删除之后用户列表中该用户消失

![此处输入图片的描述][12]


13、 命令输入错误

![此处输入图片的描述][13]


14、 错误传参

![此处输入图片的描述][14]


15、 错误日志：在data文件夹下

![此处输入图片的描述][15]


16、 基本运行情况记录在日志中：data文件夹下保存日志

![此处输入图片的描述][16]


基本使用情况如上。各条命令均能正常工作，输入命令出错之后会打印错误信息，并进行相对的提示，输入不合法的参数的时候告诉用户输入不合法，并把错误的详细信息保存在错误日志中。



  [1]: image/1.png
  [2]: image/2.png
  [3]: image/3.png
  [4]: image/createmeet.png
  [5]: image/addperson.png
  [6]: image/querymeeting.png
  [7]: image/deleteperson.png
  [8]: image/deletemeeting.png
  [9]: image/clear2meeting.png
  [10]: image/queryuser.png
  [11]: image/logout.png
  [12]: image/deleteaccount.png
  [13]: image/error.png
  [14]: image/errorllog.png
  [15]: image/errorlog.png
  [16]: image/info.png