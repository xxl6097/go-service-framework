var newModel = document.getElementById('newModel');

//https://layui.dev/docs/2/layer/#demo-more
// http://layui.xhcen.com/doc/layer.html


function init() {
    let password = localStorage.getItem('password');
    if (password) {
        checkAuth(password, (json) => {
            showMain(json)
        }, (error) => {
            //layer.msg(`认证失败 ${error}`, {icon: 0});
            showAuth()
        })
    } else {
        showAuth()
    }
    var input = document.getElementById('passwordInput')
    input.textContent = password
    input.addEventListener('keyup', function (event) {
        if (event.key === 'Enter') {
            var v = event.target.value;
            initLogin(v)
            event.target.value = ''; // Clear the input field
        }
    });
}

function showModelDialog(title,content,closeFunc) {
    var dialog = document.getElementById('dialog_id')
    var dialog_title = document.getElementById('dialog_title_id')
    dialog_title.innerText = title
    var dialog_main = document.getElementById('dialog_main_id')
    dialog_main.appendChild(content)
    dialog.style.display = "block";
    var close = document.getElementById('dialog_close')
    close.addEventListener('click', function (event) {
        event.stopPropagation();
        dialog.style.display = "none";
        closeFunc()
    });
    return dialog
}

function createCollapse(name,callback) {
    var i = document.createElement('i');
    i.className = 'layui-icon layui-icon-add-1'
    var addBtn = document.createElement('button');
    addBtn.appendChild(i)
    addBtn.className = 'layui-btn layui-btn-primary layui-btn-sm'
    addBtn.style = 'float: right;margin-top: 5px;'
    addBtn.addEventListener('click', function (event) {
        event.stopPropagation();
        layer.msg('感谢放过在下～' + name);
    });

    var i1 = document.createElement('i');
    i1.className = 'layui-icon layui-icon-delete'
    var delBtn = document.createElement('button');
    delBtn.appendChild(i1)
    delBtn.className = 'layui-btn layui-btn-primary layui-btn-sm'
    delBtn.style = 'float: right;margin-top: 5px;'
    delBtn.addEventListener('click', function (event) {
        event.stopPropagation();
        layer.msg('感谢放过在下～' + name);
    });



    var title = document.createElement('div');
    title.className = 'layui-colla-title'
    title.textContent = name
    title.appendChild(delBtn)
    title.appendChild(addBtn)



    var collapse = document.createElement('div');
    collapse.className = 'layui-collapse'
    var item = document.createElement('div');
    item.className = 'layui-colla-item'
    var content = document.createElement('div');
    content.className = 'layui-colla-content'
    var p = document.createElement('p');
    p.textContent = 'arms'
    callback(content)
    item.appendChild(title)
    item.appendChild(content)
    collapse.appendChild(item)
    return collapse
}

function creteTable(callback) {
    var table = document.createElement('table');
    var thead = document.createElement('thead');
    var tbody = document.createElement('tbody');
    var tr = document.createElement('tr');
    var th_name = document.createElement('th');
    th_name.textContent = '程序名称'
    var th_ags = document.createElement('th');
    th_ags.textContent = '运行参数'
    var th_des = document.createElement('th');
    th_des.textContent = '描述信息'
    var th_install = document.createElement('th');
    th_install.textContent = '操作'

    callback(tbody)


    tr.appendChild(th_name)
    tr.appendChild(th_ags)
    tr.appendChild(th_des)
    tr.appendChild(th_install)
    //
    // tr_body.appendChild(td_1)
    // tr_body.appendChild(td_2)
    // tr_body.appendChild(td_3)

    table.appendChild(thead)
    table.appendChild(tr)
    //tbody.appendChild(tr_body)
    table.appendChild(tbody)
    return table
}

function createNode(name,list) {
    console.log('createNode',name,list)
    node = createCollapse(name,callback=>{
        callback.appendChild(creteTable(tbody=>{
            Object.entries(list).forEach(([key, value]) => {//arm64

                var newRow = tbody.insertRow();
                var cell0 = newRow.insertCell(0);
                var cell1 = newRow.insertCell(1);
                var cell2 = newRow.insertCell(2);
                var cell3 = newRow.insertCell(3);

                var installBtn = document.createElement('button');
                installBtn.className = 'layui-btn layui-btn-xs'
                installBtn.textContent = '安装';
                installBtn.style = 'margin-right: 5px; margin-left: 5px;'
                installBtn.addEventListener('click', function () {
                    layer.msg('感谢放过在下～' + value.name);
                });

                cell0.innerText = value.name
                cell1.innerText = value.args
                cell2.innerText = value.description
                cell3.appendChild(installBtn)

            });


        }))
    })
    return node
}

function createMarket(jsonObj) {
    var div = document.createElement('div');
    Object.entries(jsonObj).forEach(([key, value]) => {
        Object.entries(value).forEach(([key1, value1]) => {
            //market.appendChild(createCollapse('windows',node))
            div.appendChild(createCollapse(key1,(callback=>{//windowns
                Object.entries(value1).forEach(([key2, value2]) => {//arm64
                   callback.appendChild(createNode(key2,value2))
                });

            })))
        });
    });
    return div
}

function onMarketHandle() {
    // var market = document.getElementById('market')
    // createMarket(market,testjson)
    // layui.element.render('collapse');
    // market.style.display = "block";

    let market = createMarket(testjson);
    var close = showModelDialog('应用市场',market,()=>{
        while (market.firstChild) {
            market.removeChild(market.firstChild);
        }
    })
    layui.element.render('collapse');
}

function onAppStoreClick() {
    var loadIndex = layer.msg('请求AppStore数据...', {
        icon: 16,
        shade: 0.01
    });
    getAppList(response=>{
        showAppStoreTable(loadIndex,response.data)
    },error=>{
        layer.close(loadIndex)
        layer.msg(error);
        onSettingAppStoreUrl('dialog')
    })
}

function onSettingAppStoreUrl(msgid) {
    if (msgid === 'dialog'){
        let dialog;
        var div = document.createElement('div');
        var div0 = document.createElement('div');
        div0.className = 'layui-form-item'
        var input = document.createElement('input');
        input.className = 'layui-input'
        input.placeholder = '请输入AppStoree地址'
        div0.appendChild(input)
        var div1 = document.createElement('div');
        div1.className = 'layui-form-item'
        var button = document.createElement('button');
        button.textContent = '确定'
        button.className = 'layui-btn'
        button.addEventListener('click', function () {
            settingAppStore(input.value,()=>{
                layer.msg('设置成功');
            },(msg)=>{
                layer.msg(msg, {icon: 0});
            })
            if(dialog){
                dialog.style.display = 'none'
            }
            while (div.firstChild) {
                div.removeChild(div.firstChild);
            }
        })
        div1.appendChild(button)
        div.appendChild(div0)
        div.appendChild(div1)
        dialog = showModelDialog('设置AppStore地址',div,()=>{
            while (div.firstChild) {
                div.removeChild(div.firstChild);
            }
        })
    }else{
    }
}

function showAppStoreTable(loadIndex,json) {
    let dialog;
    var table = creteTable(tbody=>{
        Object.entries(json).forEach(([key, value]) => {//arm64

            var newRow = tbody.insertRow();
            var cell0 = newRow.insertCell(0);
            var cell1 = newRow.insertCell(1);
            var cell2 = newRow.insertCell(2);
            var cell3 = newRow.insertCell(3);

            var installBtn = document.createElement('button');
            installBtn.className = 'layui-btn layui-btn-xs'
            installBtn.textContent = '安装';
            installBtn.style = 'margin-right: 5px; margin-left: 5px;'
            installBtn.addEventListener('click', function () {
                //layer.msg('感谢放过在下～' + value.name);
                console.log(value);
                layer.confirm(`确定按照${value.name}程序吗？`, {icon: 0}, function () {
                    newApp(value, response => {
                        console.log('sucess', response)
                        newModel.style.display = 'none';
                        layer.msg('程序新建成功', {icon: 1});

                        setTimeout(()=>{
                            getRunningApps()
                        },6000)
                        if (dialog){
                            dialog.style.display = "none";
                        }
                        clearTable(table)
                    }, err => {
                        console.log('failed', err)
                        layer.msg('程序新建失败', {icon: 0});
                        if (dialog){
                            dialog.style.display = "none";
                        }
                        clearTable(table)
                    })
                }, function () {
                    layer.msg('再会～', {icon: 1});
                    if (dialog){
                        dialog.style.display = "none";
                    }
                    clearTable(table)
                });

            });

            if (value.args){
                let arrString = value.args.join(' '); // 指定自定义分隔符
                cell1.innerText = arrString
            }else{
                cell1.innerText = '无参数'
            }
            cell0.innerText = value.name
            cell2.innerText = value.description
            cell3.appendChild(installBtn)

        });


    })
    dialog = showModelDialog('应用市场',table,()=>{
        clearTable(table)
    })
    layer.close(loadIndex)
}


function testBtn() {
    //clear()
    Object.entries(marketJson).forEach(([key2, value2]) => {//arm64
        console.log(key2,value2)
    });
}


function openNewAppDialog() {
    newModel.style.display = 'block';
}

function onCancelClick() {
    newModel.style.display = 'none';
}

function onNewAppclick() {
    let name = document.getElementById('name').value;
    let binurl = document.getElementById('binurl').value;
    let confurl = document.getElementById('confurl').value;
    let args = document.getElementById('args').value;
    let description = document.getElementById('description').value;
    let argsArray = args.match(/\S+/g); // 匹配所有非空白字符的序列
    if (name === '' || binurl === '' ) {
        layer.msg('请正确输入', {icon: 0});
        return
    }
    let jsonObj = {
        name: name,
        binUrl: binurl,
        confurl: confurl,
        args: argsArray,
        description:description
    }
    console.log(jsonObj);
    var loadIndex = layer.msg('App创建中...', {
        icon: 16,
        shade: 0.01
    });
    newApp(jsonObj, response => {
        console.log('sucess', response)
        newModel.style.display = 'none';
        layer.close(loadIndex)
        layer.msg('新建成功', {icon: 1});
        setTimeout(()=>{
            getRunningApps()
        },6000)
    }, err => {
        layer.close(loadIndex)
        console.log('failed', err)
        layer.msg('新建失败', {icon: 0});
    })
}

function clear() {
    localStorage.removeItem('password');
}

function showToast(content) {
    Toast(content, 3)
}

function Toast(content, timeout) {
    var toastElement = document.getElementById("toast");
    // 设置Toast文本
    toastElement.innerText = content;
    // 显示Toast
    toastElement.style.display = "block";
    // 3秒后隐藏Toast
    setTimeout(function () {
        toastElement.style.display = "none";
    }, 1000 * timeout);
}


function initLogin(password) {
    Login(password, response => {
        //登录成功，现实主界面，存储password
        showMain(undefined)
        localStorage.setItem('password', response.data);
        console.log('认证成功', response)
        layer.msg(`认证成功`, {icon: 0});
    }, error => {
        //showAuth()
        clear()
        console.log('failed', error)
        layer.msg(`认证失败 ${error}`, {icon: 0});
    })
}

function showAuth() {
    document.getElementById('content').style.display = 'none';
    document.getElementById('auth').style.display = 'block';
    clear()
}

function getRunningApps() {
    getAll((code, response) => {
        if (code === 200) {
            if (response.code === 0) {
                if (response.data) {
                    // 使用 for...of 循环倒序遍历数组
                    var table = document.getElementById("myTable");
                    if (table){
                        clearTable(table)
                    }
                    for (var element of response.data.reverse()) {
                        addItemByGet(element)
                    }
                } else {
                }
            }
        }
    })
}

function showMain(json) {
    document.getElementById('content').style.display = 'block';
    document.getElementById('auth').style.display = 'none';
    if (json && json !== undefined){
        document.getElementById('app_name').innerText = `${json.displayName} ${json.appVersion}`;
        document.getElementById('app_desc').innerText = json.description;
        document.title = `${json.appName} ${json.appVersion}`;
    }
    getRunningApps()
}

function insertRow(tbody, newRow, newItem) {
    var cell0 = newRow.insertCell(0);
    var cell1 = newRow.insertCell(1);
    var cell2 = newRow.insertCell(2);
    var cell3 = newRow.insertCell(3);

    // 创建按钮并设置事件处理程序
    var stopButton = NewButton('停止',`确定停止${newItem.name}程序吗？`,()=>{
        post('stop', newItem.name, (data) => {
            //showToast('停止成功')
            layer.msg('停止成功', {icon: 1});
            setTimeout(()=>{
                getRunningApps()
            },2000)
        }, (err) => {
            // showToast('停止失败')
            layer.msg('停止失败', {icon: 0});
        })
    })


    // 创建按钮并设置事件处理程序
    var restartButton = NewButton('重启',`确定重启${newItem.name}程序吗？`,()=>{
        post('restart', newItem.name, (data) => {
            layer.msg('重启成功', {icon: 1});
            setTimeout(()=>{
                getRunningApps()
            },2000)
        }, (err) => {
            layer.msg('重启失败', {icon: 0});
        })
    },'layui-btn layui-btn-normal layui-btn-xs')


    var uninstallButton = NewButton('卸载',`确定删除${newItem.name}程序吗，请慎重考虑！`,()=>{
        post('del', newItem.name, (data) => {
            layer.msg('卸载成功', {icon: 1});
            setTimeout(()=>{
                getRunningApps()
            },2000)
        }, (err) => {
            layer.msg('卸载失败', {icon: 0});
        })
    },'layui-btn layui-btn-primary layui-border-orange layui-btn-xs')

    var confButton = NewButton('配置',`确定修改${newItem.name}配置吗，请慎重考虑！`,()=>{
        postRaw('app/config', newItem.name, null,(data) => {
            layer.msg('读取成功', {icon: 1});
            console.log('读取成功',data)
            showConfigContent(`${newItem.name}的配置文件内容`,data,(value, index, elem)=>{
                console.log('修改内容',value)
                postRaw('app/config/save', newItem.name, value,(data) => {
                    layer.msg('保存成功', {icon: 1});
                }, (err) => {
                    layer.msg('读取失败', {icon: 0});
                })
            })
        }, (err) => {
            layer.msg('读取失败', {icon: 0});
        })
    },'layui-btn layui-btn-primary layui-border-red layui-btn-xs')
    var logButton = NewButton('日志',`确定修改${newItem.name}配置吗，请慎重考虑！`,()=>{
        get('read/log', newItem.name, (data) => {
            layer.msg('卸载成功', {icon: 1});
        }, (err) => {
            layer.msg('卸载失败', {icon: 0});
        })
    },'layui-btn layui-btn-primary layui-border-green layui-btn-xs')

    cell0.innerText = newItem.name
    cell1.innerHTML = newItem.status;
    cell2.appendChild(stopButton);
    cell2.appendChild(restartButton);
    cell2.appendChild(uninstallButton);
    cell2.appendChild(confButton);
    cell2.appendChild(logButton);
    cell3.innerHTML = newItem.description;
}

function showConfigContent(title,content,func) {
    layer.prompt({
        formType: 2,
        value: content,
        title: title,
        maxlength: 9999999999, // 限制最多输入500字
        area: ['500px', '300px'] // 自定义文本域宽高
    }, function (value, index, elem) {
        // alert(value);
        layer.close(index);
        func(value,index,elem)
    });
}

function NewButton(name,title,click,clsname='layui-btn layui-btn-xs') {
    var button = document.createElement('button');
    button.className = clsname;
    button.textContent = name;
    button.style = 'margin-right: 5px; margin-left: 5px;'
    button.addEventListener('click', function () {
        layer.confirm(title, {icon: 0}, function () {
            click()
        }, function () {
            layer.msg('感谢放过在下～', {icon: 1});
        });
    });
    return button
}

function clearTable(table) {
    var tbody = table.getElementsByTagName("tbody")[0];
    var tr = table.getElementsByTagName("tr")[0];
    // 移除表格主体中的所有行
    while (tbody.firstChild) {
        tbody.removeChild(tbody.firstChild);
    }
    while (tr.firstChild) {
        tr.removeChild(tr.firstChild);
    }
}

function addItemByUpload(newItem) {
    var table = document.getElementById("myTable");
    var tbody = table.getElementsByTagName("tbody")[0];
    var newRow = tbody.insertRow(0);
    insertRow(tbody, newRow, newItem)
}

function addItemByGet(newItem) {
    var table = document.getElementById("myTable");
    var tbody = table.getElementsByTagName("tbody")[0];
    var newRow = tbody.insertRow();
    insertRow(tbody, newRow, newItem)
}

function showLoading(msg) {
    // 显示loading状态
    // document.getElementById('overlay').style.display = 'block';
    document.getElementById('overlay').style.display = 'flex';
    document.getElementById('loading').innerText = msg
}

function hideLoading() {
    // 隐藏loading状态
    document.getElementById('overlay').style.display = 'none';
    document.getElementById('loading').innerText = '加载中...'
}

function isHttpOk(xhr) {
    if (xhr.status === 200 && xhr.response && xhr.response.length > 0) {
        return true;
    }
    return false
}

function getAll(callback) {
    console.log('call getall')
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/proc/getall', true);
    //xhr.open('POST', '/proc/getall', true);
    let password = localStorage.getItem('password');
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200) {
            if (xhr.readyState === 4) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj) {
                    callback(xhr.status, jsonObj)
                }
            }
        } else {
            callback(xhr.status, xhr.responseText)
        }
    };
    xhr.send();
}

function getAppConfig(callback) {
    console.log('call getall')
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/proc/app/config', true);
    //xhr.open('POST', '/proc/getall', true);
    let password = localStorage.getItem('password');
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200) {
            if (xhr.readyState === 4) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj) {
                    callback(xhr.status, jsonObj)
                }
            }
        } else {
            callback(xhr.status, xhr.responseText)
        }
    };
    xhr.send();
}

function showDialogInfo(width,height,content) {
    // 页面层
    layer.open({
        type: 1,
        area: [width, height], // 宽高 '420px', '430px'
        content: `<div class="layeropen">${content}</div>`
    });
}

function onUninstallClick() {
    var title = `确定卸载程序吗？`
    layer.confirm(title, {icon: 0}, function () {
        uninstall(()=>{
            layer.msg('程序卸载成功～');
            setTimeout(()=>{
                getRunningApps()
            },2000)
        },()=>{
            layer.msg('程序卸载成功～');
        })
    }, function () {
        layer.msg('感谢放过在下～', {icon: 1});
    });


}

function getDeviceInfo() {
    const url = `/proc/info`;
    console.log(url)
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    let password = localStorage.getItem('password');
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        if (xhr.status === 200) {
            if (xhr.readyState === 4) {
                console.log('====',xhr.readyState,xhr.status,xhr.response)
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj) {
                    //layer.msg(`${JSON.stringify(jsonObj.data)}`, {icon: 0});
                    showDialogInfo('420px', '730px',JSON.stringify(jsonObj.data))
                }
            }
        } else {
        }
    };
    xhr.send();
}

function get(path, name, sucess, failed) {
    const url = `/proc/${path}?name=${name}`;
    console.log(url)
    var xhr = new XMLHttpRequest();
    xhr.open('GET', url, true);
    xhr.responseType = 'blob'; // 因为我们要处理二进制文件
    let password = localStorage.getItem('password');
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200) {
            if (xhr.readyState === 4) {
                // jsonObj = JSON.parse(xhr.response)
                // if (jsonObj) {
                //     sucess(jsonObj)
                // }
                // 请求成功，创建一个链接来下载文件
                // var url = window.URL.createObjectURL(xhr.response);
                // var a = document.createElement('a');
                // a.href = url;
                // a.download = 'file.txt'; // 下载时文件的名称
                // document.body.appendChild(a);
                // a.click();
                // a.remove(); // 清理
                sucess(xhr)
            }
        } else {
            failed(xhr.responseText)
        }
    };
    xhr.send();
}

function post(path, name, sucess, failed) {
    const url = `/proc/${path}?name=${name}`;
    console.log(url)
    var xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    let password = localStorage.getItem('password');
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200) {
            if (xhr.readyState === 4) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj) {
                    sucess(jsonObj)
                }
            }
        } else {
            failed(xhr.responseText)
        }
    };
    xhr.send();
}

function postRaw(path, name,value, sucess, failed) {
    const url = `/proc/${path}?name=${name}`;
    console.log(url)
    var xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    let password = localStorage.getItem('password');
    xhr.setRequestHeader("accessToken", password)
    // 设置请求头，这里假设我们发送的是纯文本
    xhr.setRequestHeader('Content-Type', 'text/plain');
    xhr.onreadystatechange = function () {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200) {
            if (xhr.readyState === 4) {
                sucess(xhr.response)
            }
        } else {
            failed(xhr.responseText)
        }
    };
    if (value){
        xhr.send(value);
    }else{
        xhr.send();
    }

}

function getAppList(sucess,failed) {
    console.log('call getall')
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/proc/app/list', true);
    //xhr.open('POST', '/proc/getall', true);
    let password = localStorage.getItem('password');
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            console.log('getAppList====', xhr.readyState, xhr.status,xhr.response)
            if (xhr.status === 200) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj && jsonObj.code === 0) {
                    sucess(jsonObj)
                } else {
                    failed(jsonObj.msg)
                }
            } else {
                failed(xhr.status)
            }
        }
    };
    xhr.send();
}
function Login(password, sucess, failed) {
    const url = `/login`;
    console.log(url)
    var xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            console.log('Login====', xhr.readyState, xhr.status,xhr.response)
            if (xhr.status === 200) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj && jsonObj.code === 0) {
                    sucess(jsonObj)
                } else {
                    failed(jsonObj.msg)
                }
            } else {
                failed(xhr.status)
            }
        }
    };
    xhr.send();
}

function checkAuth(password, sucess, failed) {
    const url = `/auth`;
    console.log(url)
    var xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            console.log('checkAuth====', xhr.readyState, xhr.status,xhr.response)
            if (xhr.status === 200) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj && jsonObj.code === 0) {
                    sucess(jsonObj.data)
                } else {
                    failed(jsonObj.msg)
                }
            } else {
                failed(xhr.status)
            }
        }
    };
    xhr.send();
}

function uninstall(sucess,failed) {
    const url = `/uninstall`;
    console.log(url)
    let password = localStorage.getItem('password');
    var xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            console.log('checkAuth====', xhr.readyState, xhr.status,xhr.response)
            if (xhr.status === 200) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj && jsonObj.code === 0) {
                    sucess()
                } else {
                    failed(jsonObj.msg)
                }
            } else {
                failed(xhr.status)
            }
        }
    };
    xhr.send();
}

function settingAppStore(appsurl,sucess,failed) {
    const url = `/setting/appstore?url=${appsurl}`;
    console.log(url)
    let password = localStorage.getItem('password');
    var xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("accessToken", password)
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 4) {
            console.log('settingAppStore====', xhr.readyState, xhr.status,xhr.response)
            if (xhr.status === 200) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj && jsonObj.code === 0) {
                    sucess()
                } else {
                    failed(jsonObj.msg)
                }
            } else {
                failed(xhr.status)
            }
        }
    };
    xhr.send();
}

function newApp(jsonObject, sucess, failed) {
    console.log('call newApp')
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/proc/new', true);
    let password = localStorage.getItem('password');
    xhr.setRequestHeader("accessToken", password)
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.onreadystatechange = function () {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200) {
            if (xhr.readyState === 4) {
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj) {
                    sucess(jsonObj)
                }
            }
        } else {
            failed(xhr.responseText)
        }
    };
    // 序列化 JavaScript 对象为 JSON 字符串
    var jsonString = JSON.stringify(jsonObject);
    xhr.send(jsonString);
}

testjson = [
    {
        "windows": {
            "arm64": [
                {
                    "name": "frpc",
                    "args": [
                        "-c",
                        "frpc.toml"
                    ],
                    "description": "frp测试描述信息"
                },
                {
                    "name": "wechat",
                    "args": [
                        "-d",
                        "conf.toml"
                    ],
                    "description": "微信应用程序，用于测试"
                }
            ],
            "amd64": [
                {
                    "name": "frpc",
                    "args": [
                        "-c",
                        "frpc.toml"
                    ],
                    "description": "frp测试描述信息"
                },
                {
                    "name": "QQ",
                    "args": [
                        "-d",
                        "qq.toml"
                    ],
                    "description": "QQ应用程序，用于测试"
                }
            ]
        }
    },
    {
        "linux": {
            "arm64": [
                {
                    "name": "frpc",
                    "args": [
                        "-c",
                        "frpc.toml"
                    ],
                    "description": "frp测试描述信息"
                },
                {
                    "name": "dingtalk",
                    "args": [
                        "-d",
                        "dingtalk.toml"
                    ],
                    "description": "dingtalk应用程序，用于测试"
                }
            ],
            "amd64": [
                {
                    "name": "frpc",
                    "args": [
                        "-c",
                        "frpc.toml"
                    ],
                    "description": "frp测试描述信息"
                },
                {
                    "name": "surge",
                    "args": [
                        "-d",
                        "config.toml"
                    ],
                    "description": "surge应用程序，用于测试"
                }
            ]
        }
    }
]

marketJson = {
    "windows": {
        "arm64": [
            {
                "name": "frpc",
                "args": [
                    "-c",
                    "frpc.toml"
                ],
                "description": "frp测试描述信息"
            },
            {
                "name": "wechat",
                "args": [
                    "-d",
                    "conf.toml"
                ],
                "description": "微信应用程序，用于测试"
            }
        ],
        "amd64": [
            {
                "name": "frpc",
                "args": [
                    "-c",
                    "frpc.toml"
                ],
                "description": "frp测试描述信息"
            },
            {
                "name": "QQ",
                "args": [
                    "-d",
                    "qq.toml"
                ],
                "description": "QQ应用程序，用于测试"
            }
        ]
    },
    "linux": {
        "arm64": [
            {
                "name": "frpc",
                "args": [
                    "-c",
                    "frpc.toml"
                ],
                "description": "frp测试描述信息"
            },
            {
                "name": "dingtalk",
                "args": [
                    "-d",
                    "dingtalk.toml"
                ],
                "description": "dingtalk应用程序，用于测试"
            }
        ],
        "amd64": [
            {
                "name": "frpc",
                "args": [
                    "-c",
                    "frpc.toml"
                ],
                "description": "frp测试描述信息"
            },
            {
                "name": "surge",
                "args": [
                    "-d",
                    "config.toml"
                ],
                "description": "surge应用程序，用于测试"
            }
        ]
    },
    "darwin": {
        "arm64": [
            {
                "name": "frpc",
                "args": [
                    "-c",
                    "frpc.toml"
                ],
                "description": "frp测试描述信息"
            },
            {
                "name": "dingtalk",
                "args": [
                    "-d",
                    "dingtalk.toml"
                ],
                "description": "dingtalk应用程序，用于测试"
            }
        ],
        "amd64": [
            {
                "name": "frpc",
                "args": [
                    "-c",
                    "frpc.toml"
                ],
                "description": "frp测试描述信息"
            },
            {
                "name": "surge",
                "args": [
                    "-d",
                    "config.toml"
                ],
                "description": "surge应用程序，用于测试"
            }
        ]
    }
}

init()