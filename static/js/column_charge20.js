let create_column = document.querySelector('.create_column');
let column_title = document.querySelector('.column_title');
let column_introduction = document.querySelector('.column_introduction');
let button = document.querySelector('.r-side').querySelector('button');
let but = document.querySelector('.create_column').querySelector('button');
let cover = document.querySelector('.create_column').querySelector('input');
let column_list = document.querySelector('.column_list')
let bg = document.querySelector('.bg')
let photo = ' '
let date = new Date();
let year = date.getFullYear();
let month = date.getMonth() + 1;
let dates = date.getDate();
let hour = date.getHours();
let minite = date.getMinutes();
let second = date.getSeconds();
let time = year + '-' + month + '-' + dates + ' ' + hour + ':' + minite + ':' + second;

cover.onchange = function (evt) {
    let reader = new FileReader();
    reader.readAsDataURL(evt.target.files[0]);
    reader.onload = function (e) {
        photo = e.target.result
    }
}

async function ask_for_column() {
    let obj={
        user_id:localStorage.getItem('myself_id')
    }
    let res = await fetch('http://localhost:8080/getcolumns',{
        method: 'post',
        headers: {
            'content-type': 'application/json'
        },
        body:JSON.stringify(obj)
    });

    let res2 = await res.json();
    let result = res2.msg;
    for (var i = 0; i < result.length; i++) {
        let div = document.createElement('div');
        column_list.appendChild(div)
        let img = document.createElement('img')
        let h2 = document.createElement('h2');
        let span = document.createElement('span');
        let del = document.createElement('button');
        div.appendChild(img)
        div.appendChild(h2)
        div.appendChild(span)
        div.appendChild(del)
        img.src = result[i].cover
        h2.innerHTML = result[i].column_title
        span.innerHTML = result[i].column_intruduction
        del.innerHTML='删除'
        div.id = result[i].id
        h2.id = result[i].id
        h2.onclick = function (e) {
            location.href = '/column?id=' + e.target.id
        }
        div.onclick = function (e) {
            console.log("4")
        }

        let d = result[i].id
        del.onclick=async function(){
            let obj={
                column_id:d
            }
            let res = await fetch('http://localhost:8080/deletecolumn',{
                method:'post',
                headers:{
                    'Content-type':'application/json'
                },
                body:JSON.stringify(obj)
            });
            let res2 = await res.json();
            if(res2.msg=='删除专栏'){
                column_list.removeChild(div)
            }
        }

    }
}
ask_for_column()



button.onclick = function () {
    create_column.style.display = 'block'
    bg.style.display = 'block'
}
but.onclick = async function () {
    let obj = {
        column_title: column_title.value,
        column_intruduction: column_introduction.value,
        phoneoremail: localStorage.getItem('phoneoremail'),
        cover: photo,
        time:time
    }
    let res = await fetch('http://localhost:8080/create_column', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    let res2 =await res.json()
    let result = res2.msg
    if (result == '创建成功') {
        column_title.innerHTML = ' '
        column_introduction.innerHTML = ''
        create_column.style.display = 'none'
        bg.style.display = 'none'
        let div = document.createElement('div');
        column_list.appendChild(div)
        let img = document.createElement('img')
        let h2 = document.createElement('h2');
        let span = document.createElement('span');
        let del = document.createElement('button');
        div.appendChild(img)
        div.appendChild(h2)
        div.appendChild(span)
        div.appendChild(del)
        img.src=photo
        h2.innerHTML=column_title.value
        span.innerHTML=column_introduction.value,
            del.innerHTML='删除'
        let d=res2.id
            del.onclick = async function () {
                let obj = {
                    column_id: d
                }
                let res = await fetch('http://localhost:8080/deletecolumn', {
                    method: 'post',
                    headers: {
                        'Content-type': 'application/json'
                    },
                    body: JSON.stringify(obj)
                });
                let res2 = await res.json();
                if (res2.msg == '删除专栏') {
                    column_list.removeChild(this.parentNode)
                    // div.remove()
                }
            }
    }
}



