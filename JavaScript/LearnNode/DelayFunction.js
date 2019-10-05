function delay(){
  setTimeout(function(){
    console.log('delay 1 second');
  },2000);
  return;
}
function getAll(){
  var arr =[
    "f1",
    "f2",
    "f3"
  ]
  for(var i=0;i<arr.length;i++){
    delay();
    console.log(arr[i]);
  }
}
getAll();