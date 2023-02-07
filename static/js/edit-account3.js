let ret = document.querySelector('.ret')
ret.onclick = function () {
    location.href = '/user/home'
}
let li = document.querySelector('.side').querySelectorAll('li')
li[0].onclick = function () {
    location.href = '/user/settings/profile';
    li.style.color = 'black';
    li.style.backgroundColor = '#fff';
    li[0].style.color = 'rgb(46, 139, 201)';
    li[0].style.backgroundColor = 'rgb(171, 223, 247)'
}
li[1].onclick = function () {
    alert('账号设置');
    li.style.color = 'black';
    li.style.backgroundColor = '#fff';
    li[1].style.color = 'rgb(46, 139, 201)';
    li[1].style.backgroundColor = 'rgb(171, 223, 247)';
}
for (var i = 2; i <= 4; i++) {
    li[i].onclick = function () {
        alert('暂无相关信息');
        li.style.color = 'black';
        li.style.backgroundColor = '#fff';
        li[i].style.color = 'rgb(46, 139, 201)';
        li[i].style.backgroundColor = 'rgb(171, 223, 247)'
    }
}
let ziliao = document.querySelector('.fom').querySelectorAll('input');

async function ask_for_account() {
    let res = await fetch('http://localhost:8080/account')
    let res2 = await res.json()
    let result =res2.accountinfo
    ziliao[0].value = result.phone
    ziliao[1].value = result.weixin
    ziliao[2].value = result.xinlang
    ziliao[3].value = result.github
}
ask_for_account()

let but = document.querySelector('button');
but.onclick =async function () {
    let obj = {
        phone: ziliao[0].value,
        weixin: ziliao[1].value,
        xinlang: ziliao[2].value,
        github: ziliao[3].value,
    }
    let rel =await fetch('http://localhost:8080/user/settings/account', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    location.href='/user'

}