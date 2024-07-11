var newModel = document.getElementById('newModel');


function init() {
    let password = localStorage.getItem('password');
    if (password) {
        checkAuth(password, () => {
            showMain()
        }, (error) => {
            layer.msg(`认证失败 ${error}`, {icon: 0});
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


function testBtn() {
    clear()
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
    if (name === '' || binurl === '' || confurl === '' || args === '') {
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
    newApp(jsonObj, response => {
        console.log('sucess', response)
        newModel.style.display = 'none';
        layer.msg('新建成功', {icon: 1});
    }, err => {
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
        showMain()
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

function showMain() {
    document.getElementById('content').style.display = 'block';
    document.getElementById('auth').style.display = 'none';
    getAll((code, response) => {
        if (code === 200) {
            if (response.code === 0) {
                if (response.data) {
                    // 使用 for...of 循环倒序遍历数组
                    for (var element of response.data.reverse()) {
                        addItemByGet(element)
                    }
                } else {
                }
            }
        }
    })
}

function insertRow(tbody, newRow, newItem) {
    var cell0 = newRow.insertCell(0);
    var cell1 = newRow.insertCell(1);
    var cell2 = newRow.insertCell(2);
    var cell3 = newRow.insertCell(3);

    var stopBtn = document.createElement('button');
    stopBtn.className = 'layui-btn layui-btn-xs'
    stopBtn.style = 'margin-right: 5px; margin-left: 5px;'
    stopBtn.textContent = '停止';
    stopBtn.addEventListener('click', function () {
        var title = `确定停止${newItem.name}程序吗？`
        layer.confirm(title, {icon: 0}, function () {
            post('stop', newItem.name, (data) => {
                //showToast('停止成功')
                layer.msg('停止成功', {icon: 1});
            }, (err) => {
                // showToast('停止失败')
                layer.msg('停止失败', {icon: 0});
            })
        }, function () {
            layer.msg('感谢放过在下～', {icon: 1});
        });

    });


    // 创建按钮并设置事件处理程序
    var restartBtn = document.createElement('button');
    restartBtn.className = 'layui-btn layui-btn-xs'
    restartBtn.textContent = '重启';
    restartBtn.style = 'margin-right: 5px; margin-left: 5px;'
    restartBtn.addEventListener('click', function () {
        var title = `确定重启${newItem.name}程序吗？`
        layer.confirm(title, {icon: 0}, function () {
            post('restart', newItem.name, (data) => {
                layer.msg('重启成功', {icon: 1});
            }, (err) => {
                layer.msg('重启失败', {icon: 0});
            })
        }, function () {
            layer.msg('感谢放过在下～', {icon: 1});
        });
    });

    var deleteBtn = document.createElement('button');
    deleteBtn.className = 'layui-btn layui-btn-xs'
    deleteBtn.textContent = '卸载';
    deleteBtn.style = 'margin-right: 5px; margin-left: 5px;'
    deleteBtn.addEventListener('click', function () {
        var title = `确定删除${newItem.name}程序吗，请慎重考虑！`
        layer.confirm(title, {icon: 0}, function () {
            post('del', newItem.name, (data) => {
                layer.msg('删除成功', {icon: 1});
            }, (err) => {
                layer.msg('删除失败', {icon: 0});
            })
        }, function () {
            layer.msg('感谢放过在下～', {icon: 1});
        });
    });

    cell0.innerText = newItem.name
    cell1.innerHTML = newItem.status;
    cell2.appendChild(stopBtn);
    cell2.appendChild(restartBtn);
    cell2.appendChild(deleteBtn);
    cell3.innerHTML = newItem.description;
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

init()