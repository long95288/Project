/**
 * 以下代码在
 * https://max.book118.com
 * 控制台中执行，复制到控制台就可以了
 * !!!!需要注意的是每张图片都渲染到页面中了再运行,否则会出现部分的图片路径获取不了
 * 从控制台中复制路径字符串
 */
function getAllImageURL(){
  var allImageDivs = document.getElementsByClassName("webpreview-item");
  var urls = [];
  for(let i=0;i<allImageDivs.length;i++){
    var url = allImageDivs[i].children[0].src;
    // 打印每张图片的url
    console.log(i+":"+url);
    urls.push(url);
  }
  localStorage.setItem("urls",urls);
}
(function(){
  getAllImageURL();
  console.log("==以下内容是图片的路径字符串===");
  console.log(localStorage.getItem("urls"));
  console.log("=============================");
})();
//===================================================================


/**
 * 以下代码在 
 * https://view-cache.book118.com/
 * 的控制台中执行。
 * !!!! 执行前需要将获得的url字符串粘贴到list中
 */
window.downloadImage = function (url,filename) {
  // 请求延迟,最好设置大一点 64000ms
  var timeout = Math.random() * 64000;
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
        downloadElement.download = filename; //下载后文件名
        document.body.appendChild(downloadElement);
        downloadElement.click(); //点击下载
        document.body.removeChild(downloadElement); //下载完成移除元素
        window.URL.revokeObjectURL(href); //释放掉blob对象
      }
    }
    xhr.send()
  }, timeout);
}
function getAllImage(){
  // 将路径url字符串粘贴到list中
  var list = "https://view-cache.book118.com/view2/M01/16/34/wKh2BVyH4RyAcnouAAEHpAyOcoY928.png,https://view-cache.book118.com/view2/M02/16/34/wKh2BVyH4RyANzphAAHL4eceLiw186.png,https://view-cache.book118.com/view3/M02/16/32/wKh2BVyH4R2AYG5EAAI89TpsVHk269.png,https://view-cache.book118.com/view1/M00/16/31/wKh2BVyH4R6AX9bdAADlB_biex4319.png,https://view-cache.book118.com/view1/M00/16/30/wKh2BVyH4RyAIbaTAAEUpMTqGWg655.png,https://view-cache.book118.com/view1/M01/16/30/wKh2BVyH4RyAMlAUAAEpL3t2xZs253.png,https://view-cache.book118.com/view1/M03/16/30/wKh2BVyH4RyATu-bAAEFDDRJGJY820.png,https://view-cache.book118.com/view2/M05/16/34/wKh2BVyH4RyAJ8V9AAFxw9gWkGU106.png,https://view-cache.book118.com/view1/M02/16/30/wKh2BVyH4RyAe9RoAAEKqhNGG5M602.png,https://view-cache.book118.com/view3/M05/16/32/wKh2BVyH4R6AcIDGAAEfrwdj5Vk739.png,https://view-cache.book118.com/view1/M02/16/30/wKh2BVyH4RyATiisAADSsJ--l2w421.png,https://view-cache.book118.com/view3/M04/16/31/wKh2BVyH4RyAVhjmAAEtQn8NYw0776.png,https://view-cache.book118.com/view2/M00/07/00/wKh2BVyf-raAVwxqAAEguLAAFg8611.png,https://view-cache.book118.com/view1/M05/07/04/wKh2BVyf-q2AfdkXAAENcgey7YI359.png,https://view-cache.book118.com/view3/M04/07/0B/wKh2BVyf-qqASZ_YAAEYnvBJfSM624.png,https://view-cache.book118.com/view2/M05/07/00/wKh2BVyf-qqAbcXFAAEa3w24bqk722.png,https://view-cache.book118.com/view1/M00/07/04/wKh2BVyf-q6AZWpFAADAkLgty-0180.png,https://view-cache.book118.com/view1/M04/07/04/wKh2BVyf-qWACHVlAAFQno46No4355.png,https://view-cache.book118.com/view1/M05/07/04/wKh2BVyf-rKAGKNTAAEG-fDW_iE580.png,https://view-cache.book118.com/view2/M05/07/00/wKh2BVyf-rqAW0VtAAD6WbG9XAE017.png,https://view-cache.book118.com/view1/M00/07/04/wKh2BVyf-rWAOsGFAAD-fxooyi4588.png,https://view-cache.book118.com/view2/M00/07/00/wKh2BVyf-r2Ac7vLAADgj60LuHQ640.png,https://view-cache.book118.com/view1/M04/07/04/wKh2BVyf-rKAcfpYAAGTas-Z8dA643.png,https://view-cache.book118.com/view3/M03/07/0B/wKh2BVyf-reAM9UJAAF_pDlaIws585.png,https://view-cache.book118.com/view1/M01/07/04/wKh2BVyf-seAW90_AAD8QzpBYOw846.png,https://view-cache.book118.com/view1/M02/07/04/wKh2BVyf-sOAUOmvAAEYFXltr7s833.png,https://view-cache.book118.com/view1/M04/07/04/wKh2BVyf-suAXeJMAAERooNZQR0109.png,https://view-cache.book118.com/view2/M00/07/00/wKh2BVyf-tGAPD5XAADds7QvoEw290.png,https://view-cache.book118.com/view3/M02/07/0B/wKh2BVyf-tCAMTm6AAFHgfqI8Bo509.png,https://view-cache.book118.com/view3/M01/07/0B/wKh2BVyf-s-AXzwbAADiohrFQVY811.png,https://view-cache.book118.com/view1/M00/07/04/wKh2BVyf-tiAFnfAAAFDHvjZY1o835.png,https://view-cache.book118.com/view1/M01/07/04/wKh2BVyf-tmAYIefAADkn0QPgrM409.png,https://view-cache.book118.com/view2/M01/07/00/wKh2BVyf-tqAEv1yAAEDSVa_fAI050.png,https://view-cache.book118.com/view3/M03/07/0B/wKh2BVyf-tqAfneoAADvCDRyUOg258.png,https://view-cache.book118.com/view1/M02/07/04/wKh2BVyf-tqAHviBAAEqOnR-49Q466.png,https://view-cache.book118.com/view2/M02/07/00/wKh2BVyf-tuACSgqAADf9KD4gms433.png,https://view-cache.book118.com/view4/M00/22/3F/wKh2ClzaZdmAfGlmAADZLakhpSg801.png,https://view-cache.book118.com/view2/M01/15/0E/wKh2BVzaZdmADwmMAAFow0vdnaw663.png,https://view-cache.book118.com/view3/M00/14/14/wKh2BVzaZdmARpvoAABJGwg9c8c975.png,https://view-cache.book118.com/view4/M04/22/3E/wKh2CVzaZdmAYir9AADztojfNGo864.png,https://view-cache.book118.com/view2/M03/15/0E/wKh2BVzaZdmAe1N-AAE_5RaxTC4552.png,https://view-cache.book118.com/view5/M02/23/09/wKh2CVzaZdmAVjzsAAE00L2Ggk4457.png,https://view-cache.book118.com/view6/M03/22/25/wKh2BFzaZU6AYXXnAAC8x8tbt8A352.png,https://view-cache.book118.com/view1/M03/15/39/wKh2BVzaZU6AeJdjAAEvWAGwdOw342.png,https://view-cache.book118.com/view5/M01/23/01/wKh2CVzaZUuAL_xhAAD1ni9UBc8782.png,https://view-cache.book118.com/view4/M02/22/36/wKh2CVzaZUyAFY6cAAC_zH_LdKs068.png,https://view-cache.book118.com/view3/M01/14/08/wKh2BVzaZW2ALz6PAAD0HsH8Ulg586.png,https://view-cache.book118.com/view1/M02/15/39/wKh2BVzaZU6ACjr4AAC6NOPmELU149.png,https://view-cache.book118.com/view4/M02/22/37/wKh2ClzaZVaAeyJnAAEKBF1hCuw084.png,https://view-cache.book118.com/view1/M03/15/38/wKh2BVzaZUuAc9PMAAFECYyjlVg598.png,https://view-cache.book118.com/view6/M04/22/25/wKh2BFzaZU2AZL04AADpPb2bhxg039.png,https://view-cache.book118.com/view3/M02/14/05/wKh2BVzaZVOAY525AADjnH9onzo752.png,https://view-cache.book118.com/view3/M01/14/04/wKh2BVzaZUuAdKISAADAmX1RkwI049.png,https://view-cache.book118.com/view6/M00/22/25/wKh2BFzaZVCASOY5AADiBGJ5LPs943.png,https://view-cache.book118.com/view1/M02/15/3D/wKh2BVzaZYKAbSofAADh5Msjh7Y485.png,https://view-cache.book118.com/view5/M01/23/04/wKh2BFzaZYKAT5XwAADNk95vWzo706.png,https://view-cache.book118.com/view1/M02/15/3E/wKh2BVzaZYSAdLVeAACtYrQgaZw681.png,https://view-cache.book118.com/view5/M00/23/04/wKh2CVzaZYSAF5MxAADWwh0ZQaQ390.png,https://view-cache.book118.com/view1/M02/15/3E/wKh2BVzaZYWAaszTAAD6NDWVNhw035.png,https://view-cache.book118.com/view4/M04/22/3A/wKh2CVzaZYiAJ-rhAAEGwQwCCUk816.png,https://view-cache.book118.com/view2/M02/10/04/wKh2BV0nbbKAUed_AAD_qRT9Vno667.png,https://view-cache.book118.com/view2/M02/1D/23/wKh2BF0nbbmAf_PQAADsbPStaWc225.png,https://view-cache.book118.com/view4/M01/0E/3F/wKh2CV0nbbmAcgChAAEjf6ZjYdA562.png,https://view-cache.book118.com/view2/M00/10/03/wKh2BV0nbaqAFcoSAADBJzCCaSg226.png";
  var urls = list.split(",");
  for(let i=0;i<urls.length;i++){
    // 图片后缀
    var suffix = urls[i].split(".")[urls[i].split(".").length-1];
    // 图片序列名
    var filename = i+"."+suffix;
    downloadImage(urls[i],filename);
  }
}
(function(){
  getAllImage();
})();
// =======================================