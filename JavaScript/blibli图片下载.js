
function downloadImage(url) {
  console.log(url);
  setTimeout(function () {
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
  }, 2000);
}
function getImage() {
  var imageBox = document.getElementsByClassName("img-box");
  var message = "确认下载该cv的图片?"
  if (confirm(message) === true) {
    for (var i = 0; imageBox.length; i++) {
      var html = imageBox[i].innerHTML;
      // console.log(html);
      var pattern = /src=[\'\"]?([^\'\"]*)[\'\"]/i
      var srcs = pattern.exec(html)
      // console.log(srcs[0]);
      var imgSrc = "https:" + srcs[0].split("=")[1].replace("\"", "").split("@")[0];
      // 延迟1-2秒
      downloadImage(imgSrc);
    }
  }
}
function insertButton(){
  var button = document.createElement("button");
  button.setAttribute("value","下载图片");
  button.setAttribute("onclick","getImage()");
  button.innerText = "下载该cv的图片"
  document.getElementsByClassName("fixed-box")[0].append(button);
}