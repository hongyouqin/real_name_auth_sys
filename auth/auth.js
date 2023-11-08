function submitForm() {
    var idCard = document.getElementById("idCard").value;
    var fullName = document.getElementById("fullName").value;
    var resultElement = document.getElementById("result");

    // 创建一个 XMLHttpRequest 对象
    var xhr = new XMLHttpRequest();
    
    // 定义 POST 请求的URL，这里假设后端服务器运行在本地的8080端口
    var url = "http://127.0.0.1:80";
    
    // 配置请求
    xhr.open("POST", url, true);
    xhr.setRequestHeader("Content-Type", "application/x-www-form-urlencoded");
    
    // 定义请求的回调函数
    xhr.onreadystatechange = function() {
        if (xhr.readyState === 4 && xhr.status === 200) {
            // 请求成功，根据后端返回的结果显示认证结果
            resultElement.innerHTML = xhr.responseText;
        }
    };
    
    // 构建请求体
    var data = "idCard=" + idCard + "&fullName=" + fullName;
    
    // 发送请求
    xhr.send(data);
}
