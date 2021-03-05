/**
 * Created by Eric.Zhang on 2017/4/7.
 */
var joyconn_layout={};
joyconn_layout.Data={uid:""}
joyconn_layout.logout =function () {
    $.post(page_content_path+"/api/iotcomm/CommApi/AccountApi/dologout",{},function (data) {

        window.location.href = page_content_path +'/page/account/login'
    })
}
joyconn_layout.modPwd=function () {
    var html = '<div id="dialog_pwdMofify">'+ $("#pwdModifyParentDiv").html()+'</div>';

    var d1 = dialog({
        title: '修改密码',
        content: html,
        cancelValue: '取消',
        cancel: true,
        okValue :"修改",
        ok: function () {
            var curpwd = $("#dialog_pwdMofify").find("#Home_modpwd_curpwd").val();
            var newpwd = $("#dialog_pwdMofify").find("#Home_modpwd_newpwd").val();
            var repwd = $("#dialog_pwdMofify").find("#Home_modpwd_repwd").val();
            var errorNote= $("#dialog_pwdMofify").find("#pwdModifyErrorNoteDiv").find('span');
            if (curpwd == '') {
                $(errorNote).html('请输入当前密码');
                $(pwdModifyErrorNoteDiv).css("display","");
                return false;
            }
            if (newpwd == '') {
                $(errorNote).html('请输入新密码');
                $(pwdModifyErrorNoteDiv).css("display","");
                return false;
            }
            if (newpwd != repwd) {
                $(errorNote).html('两次密码输入的不一致');
                $(pwdModifyErrorNoteDiv).css("display","");
                return false;
            }
            $.ajax({
                url:page_content_path+ '/api/iotcomm/ManageApi/UserManageApi/modifyPwd',
                type: 'post',
                data: { pwd:  curpwd, npwd: newpwd },
                cache: false,
                success: function (data) {
                    if(joyconn_layout.ValidataResult(data)) {
                        if (data.code == "OperateOk") {
                            d1.close().remove();
                            dialog({
                                title: '修改成功',
                                content: '密码修改成功！',
                                quickClose: true,
                                cancel: false
                            }).showModal();
                            return;
                        }
                        else if (data.code == "ParamsError") {
                            $(errorNote).html('当前密码输入不正确');
                            $(pwdModifyErrorNoteDiv).css("display", "");
                        }
                        else {
                            $(errorNote).html('密码修改失败');
                            $(pwdModifyErrorNoteDiv).css("display", "");
                        }
                    }
                }
            });
        }
    });
    d1.showModal();
}
//ajax结果返回值验证，结果为空或失去登录或访问没有权限都会返回false，内部已做好弹出层提示
/**
 * 请求结果验证
 * @param data 服务器返回的数据
 * @param extNote 是否对除未登录、权限不足以外的验证做出提示，true表示自己进行提示，不传或false表示系统自动提示
 * @returns {boolean} true验证通过 false验证失败
 * @constructor
 */
joyconn_layout.ValidataResult=function (data,extNote) {
    if(!data||data==""){
        return false;
    }else if(data.code=="NoRule"){
        var errmsg = "";
        if(data.errorMsg){
           try{
                var msgs = JSON.parse(data.errorMsg);
                if(msgs){                    
                    if(msgs.length>0){                        
                        errmsg += "<br />缺少如下权限：";
                        $(msgs).each(function(i,msg){
                            if(msg.length>0){
                                errmsg += "<br />";
                                $(msg).each(function(j,l){
                                    errmsg += l +" > ";
                                })
                            }
                            errmsg = errmsg.substr(0,errmsg.length-3);                     
                        })
                    }
                }
           }catch(e){
                errmsg = data.errorMsg
           }
        }
        dialog({
            title: '访问失败',
            content: '您没有访问该数据的权限！'+errmsg,
            quickClose: true,
            cancel: false}).showModal();
        return false;

    }else if(data.code=='NoLogin'||data.code=='LoginFail'){
        dialog({
            title: '访问失败',
            content: '您已失去登录状态！请重新登录系统！',
            cancel: false,
            okValue:'去登陆',
            ok:function () {
                window.location.href= page_content_path+'/page/account/login?ref='+encodeURIComponent(window.location.href);
            }
        }).showModal();
        return false;
    }else if(data.code=="ServiceError"&&!extNote){
        dialog({
            title: '访问失败',
            content: '服务器暂时繁忙，请联系工作人员！',
            quickClose: true,
            cancel: false}).showModal();
        return false;

    }else if(data.code=="ParamsError"&&!extNote){
        var errmsg="请求的参数不正确！" + (data.errorMsg?data.errorMsg:"");
        dialog({
            title: '访问失败',
            content: errmsg,
            quickClose: true,
            cancel: false}).showModal();
        return false;

    }else{
        return true;
    }
}

joyconn_layout.getInArray=function(arr,attrName,value){
    var result = null;
    if(arr){
        $(arr).each(function (index,model) {
            if(model[attrName]==value){
                result=model;
                return false;
            }
        })
    }
    return result;
}


/**
 * Created by Eric.Zhang on 2017/5/3.
 */
function getPageCount( itemCount, pageSize) {
    return Math.ceil(itemCount / pageSize);

}
//显示分页。
//{elemet:domElement，allcount:allcount，pno:1,pagesize:10}
function ShowPage(element, options, pageFuc) {
    if(options.allcount<= options.pagesize){
        return;
    }
    // $("#"+elementid).html('');
    $(element).pagination(
        {
            coping: true,
            homePage: '首页',
            endPage: '末页',
            prevContent: '上一页',
            nextContent: '下一页',
            totalData: options.allcount,
            pageCount: getPageCount(options.allcount, options.pagesize),
            showData: options.pagesize,
            current: options.pno,

            // items_per_page:options.pagesize,//	每页显示的条目数	可选参数，默认是10
            // num_display_entries:5,//	连续分页主体部分显示的分页条目数	可选参数，默认是10
            // current_page:options.pno,//	当前选中的页面	可选参数，默认是0，表示第1页
            // prev_text :"« 上一页",	//“前一页”分页按钮上显示的文字	字符串参数，可选，默认是"Prev"
            // next_text: "下一页 »",//	“下一页”分页按钮上显示的文字	字符串参数，可选，默认是"Next"
            callback: function (api) {
                pageFuc(api.getCurrent());
                // if(index>0&&index<=getPageCount(options.pagesize,options.allcount)){
                //     pageFuc(index);
                // }
                //return false;
            }
        }
    );    
}


/**
 * 初始化 tip标签鼠标悬浮事件
 * @param iElement
 */
joyconn_layout.initTipsHover=function(parent){

    parent.find('[data-toggle="tooltip"]').tooltip({container:'body'});
    var tipElements = $(parent).find(".quick_help_note_icon");
    $(tipElements).each(function (i,tipElement) {
        if($(tipElement).find('span').text()!=''){
            $(tipElement) .tooltip({container:'body',title:$(tipElement).find('span').html(),placement:"right" ,html:true});
        }

    });
}


joyconn_layout.getMe=function(){
    $.get(
        page_content_path+'/api/iotcomm/ManageApi/UserManageApi/getCurUser',function (data) {
            if(joyconn_layout.ValidataResult(data)){
                var user=data.result.basic;
                if(user){
                    $("img[tag='user_photo_tag']").attr("src",user.p_imageurl);
                    $("span[tag='user_nick_tag']").html(user.p_username);
                    $("span[tag='user_phone_tag']").html(user.p_phonenumber);
                }
            }

        }
    );
}
