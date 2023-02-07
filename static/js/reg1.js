let ac = document.getElementById("ac");
let pw = document.getElementById("pw");
let pw2 = document.getElementById("pw2");
let button = document.getElementById("but");
button.onclick = async function () {
    let account = ac.value;
    let password = pw.value;
    let password2 = pw2.value;
    if (password.length < 4 || password.length > 10 || password != password2) {
        alert("注册失败，请重试！");
        return;
    }

    let obj = {
        phoneoremail: account,
        password: password2
    }
    let res=await fetch('http://localhost:8080/signUp',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
        body:JSON.stringify(obj)
    })
    let result=await res.json();
    if(result.msg=='注册成功'){
        window.location.href='../login';
    }else{
        alert('已存在该用户！')
    }
}