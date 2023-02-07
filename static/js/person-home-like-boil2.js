let list=document.querySelector('.list')
let changing=document.querySelector('.likes').querySelectorAll('span')[0]
let myid = new URL(location.href).searchParams.get("id")
changing.onclick=function(){
    location.href='/user/likes?id='+myid
}
async function ask_for_like_boil() {
    let obj = {
        user_id: myid
    }
    let res = await fetch('http://localhost:8080/user/praise', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg
    for (var i = 0; i < result.length; i++) {
        let span = document.createElement('span');
        list.appendChild(span)
        let img = document.createElement('img');
        let h2 = document.createElement('h2');
        let span1 = document.createElement('span');
        let span2 = document.createElement('span');
        let span3 = document.createElement('span');
        let div1 = document.createElement('div');
        let div2 = document.createElement('div');
        let div3 = document.createElement('div');
        span.appendChild(img)
        span.appendChild(h2)
        span.appendChild(span1)
        span.appendChild(span2)
        span.appendChild(span3)
        span.appendChild(div1)
        span.appendChild(div2)
        span.appendChild(div3)
        img.src = result[i].cover
        h2.innerHTML = result[i].author
        span1.innerHTML = result[i].date
        span2.innerHTML = result[i].article_title
        span3.innerHTML = result[i].article_content
        div1.innerHTML = ''+result[i].like
        div2.innerHTML = ''+result[i].comment
        div3.innerHTML =''+ result[i].view

        div1.onclick=async function(){
            let obj={
                id:result[i].id
            }
            let res = await fetch('http://localhost:8080/like',{
                method:'post',
                headers:{
                    'Content-type':'application/json'
                },
                body:JSON.stringify(obj)
            });
            let res2 = await res.json();
            if (res2.msg == '点赞成功') {
                result[i].like++
                div1.innerHTML = '' + result[i].like

            } else {
                result[i].like--
                div1.innerHTML = '' + result[i].like
            }
        }
    }
}
ask_for_like_boil()