/**
 * Created by Eric.Zhang on 2017/4/7.
 */
var joyconn_layout_authorize = {};
joyconn_layout_authorize.getUrlParam =function(name) {
    var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)"); //构造一个含有目标参数的正则表达式对象
    var r = window.location.search.substr(1).match(reg);  //匹配目标参数
    if (r != null) return unescape(r[2]); return null; //返回参数值
}
joyconn_layout_authorize.showMenuByID = function (menuID) {
    $.get(page_content_path + '/api/system/authorize/user/permission/getMyMenuByID',{menuID:menuID},function(data){
        if (joyconn_layout.ValidataResult(data)) {
          var tree = data.data;
          if(tree){
              var root ={pId:menuID,pMenuID:menuID,pPid:0,pName:'菜单根节点',children:[]}
               joyconn_layout_authorize.convertToTree(root,tree);
               joyconn_layout_authorize.removeEmptyNode(root);
               joyconn_layout_authorize.showMenuData(root);
          }
        }
      })

}


joyconn_layout_authorize.convertToTree=function(parent,list){
    if (!list){
        return []
    }
    list.sort(function(a,b){ 
        if (a.pPid!=b.pPid) {
            return a.pPid-b.pPid
        }else{
            return a.pSort-b.pSort
        }
    })
    $(list).each(function(i,obj){
        if(obj){
            if(obj.pPid==parent.pId){
                if(!parent.children){
                    parent.children=[]
                }
                parent.children.push(obj)
            }else{
                $(list).each(function(j,obj2){
                    if(obj2){
                        if(obj.pPid==obj2.pId){
                            if(!obj2.children){
                                obj2.children=[]
                            }
                            if($.inArray(obj,obj2.children)<0){
                                obj2.children.push(obj)
                            }                            
                        }
                    }
                })
            }
        }
    })

}
joyconn_layout_authorize.removeEmptyNode=function(list){
    if(list!=null&&list.length>0){
        var childModel;
        for(var i=0;i<list.length;){
            childModel= list[i];
            if(childModel.pType==1){
                 removeEmptyNode(childModel.children);
            }
            if(childModel.pType==2||(childModel.children!=null&&childModel.children.length>0)){
                i++;
            }else{
                delete childModel
            }
        }

    }

}
joyconn_layout_authorize.showMenuData=function(dataTree){
    var curPageUrl = window.location.pathname;

    function fillMenu(container, model) {
        var pUrl = model.pUrl;
        if(pUrl.indexOf('javascript:')!=0&&pUrl.indexOf('http://')!=0&&pUrl.indexOf('https://')!=0){
            pUrl =page_content_path+pUrl
        }
        var modelUrl =pUrl;
        if (modelUrl.indexOf(":") > 0) {
            modelUrl = modelUrl.substr(0, modelUrl.indexOf(":"));
        }
        if (modelUrl.pParams) {
            if (modelUrl.indexOf("?") > 0) {
                modelUrl += "&"+modelUrl.pParams;
            }else{
                modelUrl += "?"+modelUrl.pParams;
            }
        }

        if (model.pMenuID != 0) {
            var _element = $(' <li class="nav-item"><a class="nav-link" href="#"> <i class="nav-icon  ' + ((!model.pIcon || model.pIcon == "#") ? 'far fa-circle' : model.pIcon) + '"></i><p> ' + model.pName + '</p></a></li>');
            if (model.children && model.children.length > 0) {
                _element.find('p:eq(0)').append(' <i class="fas fa-angle-left right"></i>')
                var ul = $('<ul class="nav nav-treeview"></ul>');
                $(_element).append(ul);
                $(model.children).each(function (index, _model) {
                    fillMenu(ul, _model);
                });
            } else {
                _element.find('a').attr('href', modelUrl)
            }
            $(container).append(_element);
            // if (model.pLevel == 1) {
               
            // } else if (model.pLevel == 2) {
            //     var li = $('<li  class="nav-item"><a  class="nav-link" href="' + modelUrl + ' "><i class="nav-icon  ' + ((!model.pIcon || model.pIcon == "#") ? 'far fa-circle' : model.pIcon) + '"></i>' + model.pName + '</a></li>');
            //     $(container).append(li);
            //     if (model.pUrl.toLowerCase() == curPageUrl.toLowerCase()) {
            //         // $(container).parent().addClass("active");
            //         // $(li).addClass("active");
            //     }
            // }
        }
    }

    // var curInMenu=false;
    if (dataTree.children && dataTree.children.length > 0) {

        $(dataTree.children).each(function (index, model) {
            fillMenu($("#MAIN_NAVIGATION_MENU"), model);
        });

        // var curPageUrl = window.location.pathname;
        // var active_ele = $(".menu_node").find("a[href='" + curPageUrl + "']");
        // $(active_ele).addClass("active");
        // $(active_ele).parents(".menu_node").addClass("active");
    }
    //  $("#MAIN_NAVIGATION_MENU").after(_html);
}
/**
 *
 * @param appids 应用id数组
 * @param userid 用户id
 */
joyconn_layout_authorize.managerUserRole = function (appids, userid, appId_Names) {
    $.get( page_content_path+"/api/joyconn/authorize/PermissionUserRoleApi/getMyRoles", {
        appids: appids
    }, function (data) {
        if (!joyconn_layout.ValidataResult(data)) {
            return;
        }
        var roles = data.result;
        if (roles) {
            var div = $('<div class="box box-success user_role_select_contiar" style="border: none;box-shadow: none;overflow-y: auto; max-width: 38em;min-width: 300px;">\n' +
                '            <div class="box-body" style="border: none;box-shadow: none;">\n' +
                '            </div>\n' +
                '            <div class="box-footer" style="border: none;box-shadow: none;">\n' +
                '                <div class="form-group">\n' +
                '                    <span class="text-red"></span>\n' +
                '                </div>\n' +
                '            </div>\n' +
                '        </div>');
            //region 设置高宽
            var height = $(window).height();
            height = height - 180;
            var width = $(window).width();
            width = width - 100;
            $(div).css('width',width+'px').height(height);
            //endregion
            roles.sort((a1,a2)=>{
                if(a1.pAppid!=a2.pAppid ){
                    return a1.pAppid-a2.pAppid;
                }else{
                    return a1.pId-a2.pId 
                }
            });
            $(roles).each(function (i, a) {
                var appNameTag = "";
                if (appId_Names) {
                    appNameTag = '<span style="padding: 0 5px;">[' + appId_Names[a.pAppid] + ']</span>';
                }
                var _element = $('<div class="form-group"> ' +
                    '<label class=""  title="' + a.pName + '">'+
                    '<input type="checkbox" data-id="' + a.pId + '" data-appid="' + a.pAppid + '" value="' + a.pId + '" />' +
                    appNameTag +
                    '<span>' + a.pName + '</span>' +
                    '</label>' +
                    '</div>');
                $(div).find(".box-body").append(_element);
            });

            var userRoleModelMap = {};
            var roleLimitMap = {};
            var d = dialog({
                title: '选择角色',
                content: div,
                cancelValue: "取消",
                okValue: "确定",
                ok: function () {
                    $(div).find(".box-footer .form-group").html('<img src="/css/images/loading.gif" />');
                    var userRoleModels = [];
                    for (var roleAppid in userRoleModelMap) {
                        if (roleLimitMap[roleAppid]) {
                            userRoleModelMap[roleAppid].pRoles = JSON.stringify(roleLimitMap[roleAppid]);
                        }
                        userRoleModels.push(userRoleModelMap[roleAppid]);
                    }


                    $.post( page_content_path+"/api/joyconn/authorize/PermissionUserRoleApi/updates", {modelsJson: JSON.stringify(userRoleModels)}, function (data) {
                        if (!joyconn_layout.ValidataResult(data)) {
                            return;
                        }
                        if (data.result) {
                            d.close().remove();
                            dialog({
                                title: '成功',
                                content: '用户角色修改成功',
                                quickClose: true,
                            }).show();

                        } else {
                            $(div).find(".box-footer .form-group").html(' <span class="text-red" >用户角色修改失败</span>');
                        }

                    });
                }
            }).showModal();
            $(div).find('.box-body input').iCheck({
                checkboxClass: 'icheckbox_square-blue',
                radioClass: 'iradio_square-blue',
                increaseArea: '20%' // optional
            });
            $.get( page_content_path+"/api/joyconn/authorize/PermissionUserRoleApi/selectUserRolesByAppids", {
                appids: appids,
                fUserid: userid
            }, function (data) {
                if (!joyconn_layout.ValidataResult(data)) {
                    return;
                }
                if (data.result) {
                    $(data.result).each(
                        function (i, userRoleModel) {
                            userRoleModelMap[userRoleModel.pAppid] = userRoleModel;
                            if (userRoleModel.pRoles) {
                                roleLimitMap[userRoleModel.pAppid] = JSON.parse(userRoleModel.pRoles);
                                $(roleLimitMap[userRoleModel.pAppid]).each(function (i, userRole) {
                                    $(div).find('.box-body input[data-id="' + userRole.roleID + '"]').iCheck("check");
                                })
                            }
                        }
                    )

                    $('input').on('ifChecked', function (event) {
                        var roleid = $(this).val();
                        var roleAppid = $(this).attr("data-appid");
                        if (!roleLimitMap[roleAppid]) {
                            roleLimitMap[roleAppid] = [];
                        }
                        if (!userRoleModelMap[roleAppid]) {
                            userRoleModelMap[roleAppid] = {pAppid: roleAppid, pUserID: userid, pRoles: "[]"};
                        }
                        roleLimitMap[roleAppid].push({roleID: roleid, limitDate: '2008-01-01 00:00:00'});
                    });
                    $('input').on('ifUnchecked', function (event) {
                        var roleid = $(this).val();
                        var roleAppid = $(this).attr("data-appid");
                        $(roleLimitMap[roleAppid]).each(function (i, role) {
                            if (role.roleID == roleid) {
                                roleLimitMap[roleAppid].splice(i, 1);
                                return false;
                            }
                        });
                    });
                }



            });


        } else {
            dialog({
                title: '错误',
                content: "您还没有被赋予任何角色，不能对其他用户设置角色",
                quickClose: true
            }).showModal()
        }
    });

}

joyconn_layout_authorize.showUrl=function (appid,url,method) {
    return true;
}

