/**
 * Created by Eric.Zhang on 2019/1/7.
 */
var joyconn_utils={};
 /**
   * 数组去重 
   */
joyconn_utils.unique=function (arr){
    var res = [];
    var obj = {};
    var nullValue=false;
    for(var i=0; i<arr.length; i++){
        if(arr[i]==null){
            nullValue=true;
        }else if( !obj[arr[i]] ){
            obj[arr[i]] = 1;
            res.push(arr[i]);
        }
    } 
    if(nullValue){
        res.push(null);
    }
    return res;
}

 /**
   * 将数值转换成A-Z的字母表示
   *
   * @param n 数值
   * @return 字母表示
   */
joyconn_utils.toAlphaString=function(n) {
    var ordA = 'A'.charCodeAt(0); 
	var ordZ = 'Z'.charCodeAt(0);
	var len = ordZ - ordA + 1;
	var s = "";
	while( n >= 0 ) { 
		s = String.fromCharCode(n % len + ordA) + s; 
		n = Math.floor(n / len) - 1;
	} 
	return s;
  }