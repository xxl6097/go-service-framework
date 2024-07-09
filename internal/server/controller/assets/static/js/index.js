var newModel = document.getElementById('newModel');
function testBtn(){
    clear()
}
function openNewAppDialog(){
    newModel.style.display = 'block';
}

function onCancelClick(){
    newModel.style.display = 'none';
}

function onNewAppclick(){
    let name = document.getElementById('name').value;
    let binurl = document.getElementById('binurl').value;
    let confurl = document.getElementById('confurl').value;
    let args = document.getElementById('args').value;
    let argsArray = args.match(/\S+/g); // 匹配所有非空白字符的序列
    let jsonObj = {
        name : name,
        binUrl: binurl,
        confurl: confurl,
        args: argsArray
    }
    console.log(jsonObj);
    newApp(jsonObj,response=>{
        console.log('sucess',response)
        newModel.style.display = 'none';
        showToast('新建成功')
    },err=>{
        console.log('failed',err)
        alert('失败',err)
    })
}

var authcode = ''
function init() {
    authcode = localStorage.getItem('password');
    if (authcode){
        console.log('=============init')
        login(authcode)
    }else{
        document.getElementById('content').style.display = 'none';
        document.getElementById('auth').style.display = 'block';
    }
    var input = document.getElementById('passwordInput')
    input.textContent = authcode
    input.addEventListener('keyup', function(event) {
        if (event.key === 'Enter') {
            var password = event.target.value;
            login(password)
            event.target.value = ''; // Clear the input field
        }
    });
}


function cache(){
    localStorage.setItem('key', 'value');
    var value = localStorage.getItem('key');
    localStorage.removeItem('key');
}

function clear() {
    localStorage.removeItem('password');
}

function showToast(content) {
    Toast(content,3)
}

function Toast(content,timeout) {
    var toastElement = document.getElementById("toast");
    // 设置Toast文本
    toastElement.innerText = content;
    // 显示Toast
    toastElement.style.display = "block";
    // 3秒后隐藏Toast
    setTimeout(function () {
        toastElement.style.display = "none";
    }, 1000*timeout);
}

function searchFiles(pattern) {
    var xhr = new XMLHttpRequest();
    var url = '/search';
    if (pattern){
        url += `?pattern=${pattern}`
    }
    xhr.open('GET', url, true);
    console.log('url',url);
    xhr.onreadystatechange = function () {
        if (xhr.readyState === 1) {
            // 在这里处理loading状态，例如显示loading动画
            console.log('Loading...');
            showLoading('正在获取文件清单，请稍等～')
        }else if (isHttpOk(xhr)) {//xhr.readyState === 4 &&
            filejson = JSON.parse(xhr.response)
            if (filejson.code === 0){
                console.log('searchFiles',xhr.response)
                // var table = document.getElementById("myTable");
                // var tbody = table.getElementsByTagName("tbody")[0];
                // tbody.innerHTML = '';
                clearTable()
                if (filejson.data){
                    // 使用 for...of 循环倒序遍历数组
                    for (var element of filejson.data.reverse()) {
                        addItemByGet(element)
                    }
                    showToast('搜索到'+filejson.data.length+'个结果~')
                }else{
                    showToast('未搜索到结果~')
                }
            }else{
                console.log('失败了',filejson.msg)
            }
            hideLoading()
        } else {
            // 请求失败或还未完成
            //console.error('get files err ',xhr.response);
            console.log('searchFiles files err ',xhr.readyState,xhr.status,xhr.response);
        }
    };

    xhr.send();
}

function clearTable() {
    // 获取表格对象
    var table = document.getElementById("myTable");
    // 获取表格主体
    var tbody = table.getElementsByTagName("tbody")[0];
    // 移除表格主体中的所有行
    while (tbody.firstChild) {
        tbody.removeChild(tbody.firstChild);
    }
    //getFiles()
}

function login(password) {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/auth', true);
    xhr.setRequestHeader("accessToken",password)
    xhr.onreadystatechange = function() {
        if (xhr.status === 200 && xhr.readyState === 4){
            console.log('login=>>>',xhr.readyState,xhr.status)
            document.getElementById('content').style.display = 'block';
            document.getElementById('auth').style.display = 'none';
            localStorage.setItem('password', password);
            console.log('sucess',xhr.status,xhr.responseText)
            showToast('认证成功')

            getAll((code,response) => {
                if (code === 200){
                    if (response.code === 0){
                        if (response.data){
                            // 使用 for...of 循环倒序遍历数组
                            for (var element of response.data.reverse()) {
                                addItemByGet(element)
                            }
                        }else{
                        }
                    }
                }
            })
        }else{
            console.log('failed',xhr.status)
            //showToast('认证失败')
            document.getElementById('content').style.display = 'none';
            document.getElementById('auth').style.display = 'block';
        }
    };
    xhr.send();
}

function GetConfig(callback) {
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/config', true);
    //xhr.setRequestHeader("Authorization",password)
    xhr.onreadystatechange = function() {
        console.log('====',xhr.readyState,xhr.status)
        if (isHttpOk(xhr)){
            console.log('sucess',xhr.status,xhr.responseText)
            filejson = JSON.parse(xhr.response)
            if (filejson.code === 0){
                callback(filejson.data)
            }
        }
    };
    xhr.send();
}


function insertRow(tbody,newRow,newItem) {
    var cell0 = newRow.insertCell(0);
    var cell1 = newRow.insertCell(1);
    var cell2 = newRow.insertCell(2);

    var stopBtn = document.createElement('button');
    stopBtn.className = 'layui-btn layui-btn-xs'
    stopBtn.style = 'margin-right: 5px; margin-left: 5px;'
    stopBtn.textContent = '停止';
    stopBtn.addEventListener('click', function () {
        post('stop',newItem.name,(data)=>{
            showToast('停止成功')
        },(err)=>{
            showToast('停止失败')
        })
    });


    // 创建按钮并设置事件处理程序
    var restartBtn = document.createElement('button');
    restartBtn.className = 'layui-btn layui-btn-xs'
    restartBtn.textContent = '重启';
    restartBtn.style = 'margin-right: 5px; margin-left: 5px;'
    restartBtn.addEventListener('click', function () {
        // 当按钮点击时触发的事件
        post('restart',newItem.name,(data)=>{
            showToast('停止成功')
        },(err)=>{
            showToast('停止失败')
        })
    });

    var deleteBtn = document.createElement('button');
    deleteBtn.className = 'layui-btn layui-btn-xs'
    deleteBtn.textContent = '删除';
    deleteBtn.style = 'margin-right: 5px; margin-left: 5px;'
    deleteBtn.addEventListener('click', function () {
        // 当按钮点击时触发的事件
        post('del',newItem.name,(data)=>{
            showToast('停止成功')
        },(err)=>{
            showToast('停止失败')
        })
    });

    cell0.innerText = newItem.name
    cell1.innerHTML = newItem.status;
    cell2.appendChild(stopBtn);
    cell2.appendChild(restartBtn);
    cell2.appendChild(deleteBtn);
}

function addItemByUpload(newItem) {
    var table = document.getElementById("myTable");
    var tbody = table.getElementsByTagName("tbody")[0];
    var newRow = tbody.insertRow(0);
    insertRow(tbody,newRow,newItem)
}

function addItemByGet(newItem) {
    var table = document.getElementById("myTable");
    var tbody = table.getElementsByTagName("tbody")[0];
    var newRow = tbody.insertRow();
    insertRow(tbody,newRow,newItem)
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
    if (xhr.status === 200 && xhr.response && xhr.response.length > 0){
        return true;
    }
    return false
}

function getAll(callback) {
    console.log('call getall')
    var xhr = new XMLHttpRequest();
    xhr.open('GET', '/proc/getall', true);
    //xhr.open('POST', '/proc/getall', true);
    xhr.setRequestHeader("accessToken",authcode)
    xhr.onreadystatechange = function() {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200 ){
            if (xhr.readyState === 4){
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj){
                    callback(xhr.status,jsonObj)
                }
            }
        }else{
            callback(xhr.status,xhr.responseText)
        }
    };
    xhr.send();
}

function stopApp(name,sucess,failed) {
    console.log('call stopApp')
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/proc/stop?name=' + name, true);
    xhr.setRequestHeader("accessToken",authcode)
    xhr.onreadystatechange = function() {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200 ){
            if (xhr.readyState === 4){
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj){
                    sucess(jsonObj)
                }
            }
        }else{
            failed(xhr.responseText)
        }
    };
    xhr.send();
}

function post(path,name,sucess,failed) {
    const url = `/proc/${path}?name=${name}`;
    console.log(url)
    var xhr = new XMLHttpRequest();
    xhr.open('POST', url, true);
    xhr.setRequestHeader("accessToken",authcode)
    xhr.onreadystatechange = function() {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200 ){
            if (xhr.readyState === 4){
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj){
                    sucess(jsonObj)
                }
            }
        }else{
            failed(xhr.responseText)
        }
    };
    xhr.send();
}


function newApp(jsonObject,sucess,failed) {
    console.log('call stopApp')
    var xhr = new XMLHttpRequest();
    xhr.open('POST', '/proc/new', true);
    xhr.setRequestHeader("accessToken",authcode)
    // 设置请求头
    xhr.setRequestHeader('Content-Type', 'application/json;charset=UTF-8');
    xhr.onreadystatechange = function() {
        //console.log('====',xhr.readyState,xhr.status)
        if (xhr.status === 200 ){
            if (xhr.readyState === 4){
                jsonObj = JSON.parse(xhr.response)
                if (jsonObj){
                    sucess(jsonObj)
                }
            }
        }else{
            failed(xhr.responseText)
        }
    };
    // 序列化 JavaScript 对象为 JSON 字符串
    var jsonString = JSON.stringify(jsonObject);
    xhr.send(jsonString);
}

init()