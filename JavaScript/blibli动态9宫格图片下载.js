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
  console.log("=== element =======");
  console.log(element);
  console.log("=== end ========== ");
  var button = document.createElement("button");
  button.addEventListener("click", function () {
    console.log(element);
    // 获得该box的9个图片的url
    var urlList = element.firstElementChild.firstElementChild.children;
    for (var i = 0; i < urlList.length; i++) {
      var url = "https:" + urlList[i].firstElementChild.style.backgroundImage.split("url")[1].replace(/[\(\)\"]/g, "").split("@")[0].trim();
      console.log(url);
      // console.log(timeout);
      window.downloadImage(url);
    }
  });
  button.style.backgroundColor = "#3388ff";
  button.style.color = "white";
  // button.style.height = "40px";
  button.innerText = "下载该九宫格的图片";
  console.log("按钮问题");
  element.append(button);
}

function initImageBox(){
  var allImagesBox = document.getElementsByClassName("imagesbox");
  // 每个动态插入下载图片按钮
  console.log("=======allImagesBox======");
  console.log(allImagesBox);
  console.log("=======end============");
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