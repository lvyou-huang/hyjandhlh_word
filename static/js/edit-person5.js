let ret = document.querySelector('.ret')
ret.onclick = function () {
    location.href = '/user/home'
}
let li = document.querySelector('.side').querySelectorAll('li')
li[0].onclick = function () {
    li.style.color = 'black';
    li.style.backgroundColor = '#fff';
    li[0].style.color = 'rgb(46, 139, 201)';
    li[0].style.backgroundColor = 'rgb(171, 223, 247)'
}
li[1].onclick = function () {
    location.href = '/user/settings/account';
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
let textarea = document.querySelector('textarea');
let touxiang = document.querySelector('.touxiang')
async function ask_for_person() {
    let res = await fetch('http://localhost:8080/profile')
    let res2 = await res.json()
    let result =res2.userinfo
    ziliao[0].value = result.name
    ziliao[1].value = result.position
    ziliao[2].value = result.company
    ziliao[3].value = result.web
    textarea.value = result.introduce
    touxiang.src = result.cover
}
ask_for_person()


let photo = ' '
let cover=document.querySelector('.cover')
cover.onchange = function (evt) {
    let reader = new FileReader();
    reader.readAsDataURL(evt.target.files[0]);
    reader.onload = function (e) {
        photo = e.target.result
        console.log(photo)
    }
}

let but = document.querySelector('button');
but.onclick =async function () {
    let obj = {
        name: ziliao[0].value,
        position: ziliao[1].value,
        company: ziliao[2].value,
        web: ziliao[3].value,
        introduce: textarea.value,
        cover:photo
    }
    let rel =await fetch('http://localhost:8080/user/settings/profile', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    location.href='/user'

}
