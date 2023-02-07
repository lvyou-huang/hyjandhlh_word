let content = document.querySelector('.content')
let all = document.querySelector('.all')
let hot = document.querySelector('.hot')
async function ask_for_column() {
    let obj = {
        catalogue: '前端',
        status: ''

    }
    let res = await fetch('http://localhost:8080/course', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)

    });
    let res2 = await res.json();
    let result = res2.msg
    for (let i = 0; i < result.length; i++) {
        let div = document.createElement('div')
        let h1 = document.createElement('h1')
        let img = document.createElement('img')
        let span = document.createElement('span')
        let div1 = document.createElement('div')
        content.appendChild(div)
        div.appendChild(h1)
        div.appendChild(img)
        div.appendChild(span)
        div.appendChild(div1)
        h1.innerHTML = result[i].course_title
        img.src = result[i].cover
        span.innerHTML = result[i].abstract
        div1.innerHTML = result[i].introduction
        div.id = result[i].course_id
        div1.id = result[i].course_id
        img.id = result[i].course_id
        h1.id = result[i].course_id
        span.id = result[i].course_id
        div.onclick=function(e){
            location.href='/book?id='+e.target.id
        }

    }
}
ask_for_column()


hot.onclick = async function () {
    all.style.color = 'rgb(131, 129, 129)'
    hot.style.color = 'rgb(97, 180, 240)'
    let t = content.children.length;
    for (let i = 0; i < t; i++) {
        passage_content.removeChild(content.children[0])
    }
    let obj = {
        status: 'hot',
        catalogue: '前端'
    }
    let res = await fetch('http://localhost:8080/course', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)

    });
    let res2 = await res.json();
    let result = res2.msg
    for (let i = 0; i < result.length; i++) {
        let div = document.createElement('div')
        let h1 = document.createElement('h1')
        let img = document.createElement('img')
        let span = document.createElement('span')
        let div1 = document.createElement('div')
        content.appendChild(div)
        div.appendChild(h1)
        div.appendChild(img)
        div.appendChild(span)
        div.appendChild(div1)
        h1.innerHTML = result[i].course_title
        img.src = result[i].cover
        span.innerHTML = result[i].abstract
        div1.innerHTML = result[i].introduction
        div.id = result[i].course_id
        div1.id = result[i].course_id
        img.id = result[i].course_id
        h1.id = result[i].course_id
        span.id = result[i].course_id
        div.onclick=function(e){
            location.href='/book?id='+e.target.id
        }

    }
}


all.onclick = async function () {
    hot.style.color = 'rgb(131, 129, 129)'
    all.style.color = 'rgb(97, 180, 240)'
    let t = content.children.length;
    for (let i = 0; i < t; i++) {
        passage_content.removeChild(content.children[0])
    }
    let obj = {
        catalogue: '前端',
        status: ''

    }
    let res = await fetch('http://localhost:8080/course', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)

    });
    let res2 = await res.json();
    let result = res2.msg
    for (let i = 0; i < result.length; i++) {
        let div = document.createElement('div')
        let h1 = document.createElement('h1')
        let img = document.createElement('img')
        let span = document.createElement('span')
        let div1 = document.createElement('div')
        content.appendChild(div)
        div.appendChild(h1)
        div.appendChild(img)
        div.appendChild(span)
        div.appendChild(div1)
        h1.innerHTML = result[i].course_title
        img.src = result[i].cover
        span.innerHTML = result[i].abstract
        div1.innerHTML = result[i].introduction
        div.id = result[i].course_id
        div1.id = result[i].course_id
        img.id = result[i].course_id
        h1.id = result[i].course_id
        span.id = result[i].course_id
        div.onclick=function(e){
            location.href='/book?id='+e.target.id
        }

    }
}