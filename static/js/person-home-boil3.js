let list = document.querySelector('.list')

let myid = new URL(location.href).searchParams.get("id")
async function ask_for_boil() {
    let obj2={
        user_id:myid
    }
    let r = await fetch('http://localhost:8080/getuserinfo',{
        method:'post',
        headers:{
            'Content-type':'application/json'
        },
        body:JSON.stringify(obj2)
    });
    let r2 = await r.json();
    let photo=r2.msg.cover


    let obj = {
        user_id: myid
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
    for (var i = 0; i < result.length; i++) {
        let span = document.createElement('span');
        list.appendChild(span)
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
        
        img.src = photo
        h2.innerHTML = result[i].author
        span1.innerHTML = result[i].date
        span2.innerHTML = result[i].article_title
        span3.innerHTML = result[i].article_content
        div1.innerHTML = '点赞' + result[i].like
       

        div1.onclick = async function () {
            let obj = {
                id: result[i].id
            }
            let res = await fetch('http://localhost:8080/like', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json'
                },
                body: JSON.stringify(obj)
            });
            let res2 = await res.json();
            if (res2.msg == '点赞成功') {
                result[i].like++
                div1.innerHTML = '点赞' + result[i].like

            } else {
                result[i].like--
                div1.innerHTML = '点赞' + result[i].like
            }
        }
    }
}
ask_for_boil()