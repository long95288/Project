function getUrlList(){
  var list = []
  var imagesbox= document.getElementsByClassName("imagesbox");
  for(let i=0;i<imagesbox.length;i++){
    var urllist = imagesbox[i].firstElementChild.firstElementChild.children;
    for(let j =0;j<urllist.length;j++){
      let url = "https:"+urllist[j].firstElementChild.style.backgroundImage.split("url")[1].replace(/[\(\)\"]/g, "");
      list.push(url);
    }
  }
  localStorage.setItem("urllist",list);
}
function download(filename,content,contetnType){
  if(!contetnType){
    contetnType = "application/octet-stream";
  }
  var a = document.createElement('a');
  var blob = new Blob([content],{'type':contetnType});
  a.href = window.URL.createObjectURL(blob);
  a.download = filename
  a.click();
}

(function(){
  // getUrlList();
  // console.log(localStorage.getItem("urllist"));
  var list = localStorage.getItem("urllist").split(",");
  var content = JSON.stringify(list);
  download('urllist.json',content);
})()