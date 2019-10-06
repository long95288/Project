
function downloadImage(url) {
  console.log(url);
  return;
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
}
function getImage() {
  var imageBox = document.getElementsByClassName("img-box");
  var message = "确认下载该cv的图片?"
  if (confirm(message) === true) {
    var startNum = 0;
    var stopNum = imageBox.length;
    console.log("stopNum = "+stopNum);

    var intervalId = setInterval(function () {
      if (startNum > stopNum) {
        // 释放计时器
        clearInterval(intervalId);
        console.log("释放计时器");
        return;
      } else {
        console.log("startNum:" + startNum);
        var html = imageBox[startNum].innerHTML;
        startNum ++;
        // console.log(html);
        var pattern = /src=[\'\"]?([^\'\"]*)[\'\"]/i
        var srcs = pattern.exec(html)
        var imgSrc = "https:" + srcs[0].split("=")[1].replace(/\"/g, "").split("@")[0].trim();
        downloadImage(imgSrc);
      }
    }, 1000);
  }
}
function insertButton() {
  var button = document.createElement("button");
  button.setAttribute("value", "下载图片");
  button.setAttribute("onclick", "getImage()");
  button.style.backgroundColor = "#3388ff";
  button.style.height = "40px";
  button.style.color = "white";
  button.innerText = "下载该cv的图片"
  document.getElementsByClassName("fixed-box")[0].append(button);
}