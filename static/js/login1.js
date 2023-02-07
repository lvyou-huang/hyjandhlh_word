let ac = document.getElementById("ac");
let pw = document.getElementById("pw");
let button = document.getElementById("but");
// document.cookie="username=John Doe";
// document.cookie="phoneoremail=123";
let str = document.cookie.split('; ');
for (var i = 0; i < str.length; i++) {
    let cookie_element = str[i];
    let arr = cookie_element.split('=');
    let name = arr[0];
    if (name == 'phoneoremail') {
        // window.location.href = '../home/home2.html';
        console.log('phoneoremail')
    }
}
button.onclick = async function () {
    let obj = {
        phoneoremail: ac.value,
        password: pw.value
    }
    let res = await fetch('http://localhost:8080/login', {

        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    let result = await res.json();
    if (result.msg == '登录成功') {
        localStorage.setItem('phoneoremail', ac.value);
        localStorage.setItem('myself_id',result.user_id)
        window.location.href = '../home';
    } else {
        alert('erro')
    }
}


