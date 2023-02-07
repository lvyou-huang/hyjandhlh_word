let header = document.querySelector('.header');
let span1 = document.querySelector('.span1');
let span2 = document.querySelector('.span2');


span1.onclick = function () {
    location.href = '/user/settings/profile?id=' + myid
}
span2.onclick = function () {
    location.href = '/user/settings/profile?id=' + myid
}

var li = document.querySelector('.nav-hd').getElementsByTagName('li')
li[0].onclick = function () {
    location.href = '/user?id=' + myid
}
li[1].onclick = function () {
    location.href = '/user/posts?id=' + myid
}
li[2].onclick = function () {
    location.href = '/user/columns?id=' + myid
}
li[3].onclick = function () {
    location.href = '/user/pins?id=' + myid
}
li[4].onclick = function () {
    location.href = '/user/collections?id=' + myid
}
li[5].onclick = function () {
    location.href = '/user/following?id=' + myid
}
li[6].onclick = function () {
    location.href = '/user/likes?id=' + myid
}



let fol1 = document.querySelector('.fol1')
let fol2 = document.querySelector('.fol2')
let fol1_number = fol1.querySelector('h2')
let fol2_number = fol2.querySelector('h2')

fol1.onclick = function () {
    location.href = '/user/following?id='+myid
}
fol2.onclick = function () {
    location.href = '/user/followers?id='+myid
}

let photo = document.querySelector('.photo')
let my_name = document.querySelector('.main').querySelector('h1')


async function ask_for_following() {
    let obj = {
        user_id:myid
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
    fol1_number.innerHTML = result.length
}
ask_for_following()


async function ask_for_follower() {
    let obj = {
        user_id:myid
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
        user_id:myid
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
    my_name.innerHTML = result.phoneoremail
}
ask()