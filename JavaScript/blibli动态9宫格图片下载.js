window.downloadImage = function (url) {
  var timeout = Math.random() * 5000;
  this.setTimeout(function () {
    var xhr = new XMLHttpRequest();
    xhr.open("get", url, true);
    xhr.responseType = "blob";
    xhr.onload = function () {
      if (this.status == 200) {
        var blob = this.response;
        var downloadElement = document.createElement('a');
        var href = window.URL.createObjectURL(blob); //创建下载的链接
        downloadElement.href = href;
        downloadElement.download = name; //下载后文件名
        document.body.appendChild(downloadElement);
        downloadElement.click(); //点击下载
        document.body.removeChild(downloadElement); //下载完成移除元素
        window.URL.revokeObjectURL(href); //释放掉blob对象
      }
    }
    xhr.send()
  }, timeout);
}

window.insertButtons = function (element) {
  // 获得该动态的卡片节点的id document.getElementsByClassName("imagesbox")[0].parentElement.parentElement.parentElement.parentElement.parentElement
  var did = element.parentElement.parentElement.parentElement.parentElement.parentElement.getAttribute("data-did");
  var idListStr = this.localStorage.getItem("idList");
  // 是否是新数据
  var newData = false;
  if (idListStr === null){
    insertList = []
    this.localStorage.setItem("idList",insertList)
  }
  var idList = this.localStorage.getItem("idList").split(",");
  if (idList.indexOf(did) !== -1){
    // 已经下载过了
    newData = true;
  }

  var button = document.createElement("button");
  button.addEventListener("click", function () {
    console.log(element);
    // 判断是不是多p图片
    // document.getElementsByClassName("imagesbox")[0].firstElementChild.lastElementChild.lastElementChild.lastElementChild.children[0].firstElementChild.src
    if(element.firstElementChild.classList[0] === 'boost-wrap') {
      // 超过9p
      console.log('=========创建超过9p下载任务==========')
      let imageList = element.firstElementChild.lastElementChild.lastElementChild.lastElementChild.children;
      for(let i=0; i < imageList.length; i++) {
        let url =imageList[i].firstElementChild.src.split("@")[0].trim() ;
        console.log(`${i} url = ${url}`);
        // 下载图片
        window.downloadImage(url);
      }
      console.log('=========创建图片下载任务完成=========')
      // 添加该项目已经下载的列表中
      idList.push(did);
    } else {
      // 获得该box的9个图片的url
      var urlList = element.firstElementChild.firstElementChild.children;
      console.log('=======创建9宫格下载任务============')
      for (var i = 0; i < urlList.length; i++) {
        var url = "https:" + urlList[i].firstElementChild.style.backgroundImage.split("url")[1].replace(/[\(\)\"]/g, "").split("@")[0].trim();
        console.log(url);
        // console.log(timeout);
        window.downloadImage(url);
      }
      console.log('========创建任务完成===============')
      // 添加该项目已经下载的列表中
      idList.push(did);
    }
    // 更新缓存数据
    localStorage.setItem("idList",idList);
  });
  button.style.borderRadius="10px";
  button.style.height = "35px";
  button.style.backgroundColor="white";
  button.style.opacity= "0.5";
  button.style.borderWidth="thin";
  button.style.position="absolute";
  button.classList.add("downBtn");
  if(newData){
    button.innerText = "下载该动态的图片";
  }else{
    // 已经下载过了
    button.style.color="#02B5DA";
    button.innerText = "已经下载过了";
  }
  element.append(button);
}

function initImageBox(){
  var allImagesBox = document.getElementsByClassName("imagesbox");
  for (var i = 0; i < allImagesBox.length; i++) {
    var element22 = allImagesBox[i];
      console.log("i="+i);
    window.insertButtons(element22);
  }
  console.log("call back initImageBox")
}
(function() {
    'use strict';
    initImageBox();
    // Your code here...
})();