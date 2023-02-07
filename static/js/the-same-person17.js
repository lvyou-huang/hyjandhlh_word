let header = document.querySelector('.header');
let span1 = document.querySelector('.span1');
let span2 = document.querySelector('.span2');
let edit = document.querySelector('.edit');
let quit = document.querySelector('.quit');
let myids = new URL(location.href).searchParams.get("id")


let fo=document.querySelector('.fo')
if(myids!==localStorage.getItem('myself_id')){
    edit.style.display='none'
    quit.style.display='none'
    span1.style.display='none'
    span2.style.display='none'

    fo.style.display='block'
}else{
    edit.style.display='block'
    quit.style.display='block'
    span1.style.display='block'
    span2.style.display='block'
    fo.style.display='none'
}
let date = new Date();
let year = date.getFullYear();
let month = date.getMonth() + 1;
let dates = date.getDate();
let hour = date.getHours();
let minite = date.getMinutes();
let second = date.getSeconds();
let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;

async function notice() {
    let obj = {
        user_id: myids,
        date:time
    }
    let res = await fetch('http://localhost:8080/getnotice', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    if (res2.noticeornot == 0) {
        fo.innerHTML = '关注'
    } else {
        fo.innerHTML = '已关注'
    }
}
notice()
span1.onclick = function () {
    location.href = '/user/settings/profile'
}
span2.onclick = function () {
    location.href = '/user/settings/profile'
}
edit.onclick = function () {
    location.href = '/user/settings/profile'
}
quit.onclick = async function () {
    let res = await fetch('http://localhost:8080/sign_out')
    let result =await res.json();
    if (result.msg == '退出成功') {
        localStorage.removeItem('phoneoremail')
        localStorage.removeItem('myself_id')
        location.href = '/login'
    } else {
        alert('退出失败')
    }
}
var li = document.querySelector('.nav-hd').getElementsByTagName('li')
li[0].onclick = function () {
    location.href = '/user?id='+myids
}
li[1].onclick = function () {
    location.href = '/user/posts?id='+myids
}
li[2].onclick = function () {
    location.href = '/user/columns?id='+myids
}
li[3].onclick = function () {
    location.href = '/user/pins?id='+myids
}
li[4].onclick = function () {
    location.href = '/user/collections?id='+myids
}
li[5].onclick = function () {
    location.href = '/user/following?id='+myids
}
li[6].onclick = function () {
    location.href = '/user/likes?id='+myids
}



let fol1_number  = document.querySelector('.fol1')
let fol2_number  = document.querySelector('.fol2')


fol1_number .onclick = function () {
    location.href = '/user/following?id='+myids
}
fol2_number .onclick = function () {
    location.href = '/user/followers?id='+myids
}

let photo = document.querySelector('.photo')
let my_name = document.querySelector('.main').querySelector('h1')


fo.onclick=async function(){
    let obj = {
        followed_id: myids
    }
    let res = await fetch('http://localhost:8080/notice', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg
    if (result == '关注成功') {
        fo.innerHTML = '已关注'
    } else if (result == '取消关注') {
        fo.innerHTML = '关注'
    }
}


async function ask_for_following() {
    let obj = {
        user_id: myids
    }
    let res = await fetch('http://localhost:8080/user/following', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg
    let t=result.length
    fol1_number.innerHTML = t
}
ask_for_following()


async function ask_for_follower() {
    let obj = {
        user_id:myids
    }
    let res = await fetch('http://localhost:8080/user/followers', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg
    fol2_number.innerHTML = result.length
}
ask_for_follower()


async function ask() {
    let obj = {
        user_id: myids
    }
    let res = await fetch('http://localhost:8080/getuserinfo', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg
    photo.src = result.cover
    my_name.innerHTML = result.name
}
ask()