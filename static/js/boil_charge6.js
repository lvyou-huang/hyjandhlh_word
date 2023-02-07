let r_side = document.querySelector('.r-side')


async function ask_for_boil() {
    let obj = {
        user_id: localStorage.getItem('myself_id')
    }
    let res = await fetch('http://localhost:8080/user/pins', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg

    let obj2 = {
        user_id: localStorage.getItem('myself_id')
    }
    let r = await fetch('http://localhost:8080/getuserinfo', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj2)
    });
    let r2 = await r.json();
    let result2 = r2.msg

    for (var i = 0; i < result.length; i++) {
        let span = document.createElement('span');
        r_side.appendChild(span)
        let img = document.createElement('img');
        let h2 = document.createElement('h2');
        let span1 = document.createElement('span');
        let span2 = document.createElement('span');
        let span3 = document.createElement('span');
        let div1 = document.createElement('div');
        span.appendChild(img)
        span.appendChild(h2)
        span.appendChild(span1)
        span.appendChild(span2)
        span.appendChild(span3)
        span.appendChild(div1)

        img.src = result2.cover
        h2.innerHTML = result[i].author
        span1.innerHTML = result[i].date
        span2.innerHTML = ''
        span3.innerHTML = result[i].article_content
        div1.innerHTML = '点赞' + result[i].like

    }
}
ask_for_boil()


let button = document.querySelector('.r-side').querySelector('button')
let but = document.querySelector('.create_boil').querySelector('button')
let create_boil = document.querySelector('.create_boil')
let text = document.querySelector('.create_boil').querySelector('textarea')
let bg = document.querySelector('.bg')

let date = new Date();
let year = date.getFullYear();
let month = date.getMonth() + 1;
let dates = date.getDate();
let hour = date.getHours();
let minite = date.getMinutes();
let second = date.getSeconds();
let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;
button.onclick = function () {
    create_boil.style.display = 'block'
    bg.style.display = 'block'
}
but.onclick = async function () {
    let obj = {
        content: text.value,
        date: time
    }
    let res = await fetch('http://localhost:8080/pins', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    let res2 = await res.json()
    let result = res2

    create_boil.style.display = 'none'
    bg.style.display = 'none'
    text.innerHTML = ' '

    let obj2 = {
        user_id: localStorage.getItem('myself_id')
    }
    let r = await fetch('http://localhost:8080/getuserinfo', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj2)
    });
    let r2 = await r.json();
    let result2 = r2.msg




    let span = document.createElement('span');
    r_side.appendChild(span)
    let img = document.createElement('img');
    let h2 = document.createElement('h2');
    let span1 = document.createElement('span');
    let span2 = document.createElement('span');
    let span3 = document.createElement('span');
    let div1 = document.createElement('div');

    span.appendChild(img)
    span.appendChild(h2)
    span.appendChild(span1)
    span.appendChild(span2)
    span.appendChild(span3)
    span.appendChild(div1)

    console.log(result2.cover)
    img.src = result2.cover
    h2.innerHTML = result2.name
    span1.innerHTML = result.date
    span2.innerHTML = ''
    span3.innerHTML = result.content
    div1.innerHTML = '点赞' + result.like


}
