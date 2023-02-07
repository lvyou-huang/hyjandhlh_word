let list = document.querySelector('.list')

async function ask_for_data() {
    let obj2={
        user_id:localStorage.getItem('myself_id')
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
        user_id: localStorage.getItem('myself_id')
    }
    let res = await fetch('http://localhost:8080/user', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    });
    let res2 = await res.json();
    let result = res2.msg;
    for (var i = 0; i < result.length; i++) {
        if (result[i].postorboil === 0) {
            let div = document.createElement('div');
            list.appendChild(div)
            let img = document.createElement('img');
            let h2 = document.createElement('h2');
            let span1 = document.createElement('span');
            let span2 = document.createElement('span');
            let span3 = document.createElement('span');
            let div1 = document.createElement('div');
            let div2 = document.createElement('div');
            let div3 = document.createElement('div');
            div.appendChild(img)
            div.appendChild(h2)
            div.appendChild(span1)
            div.appendChild(span2)
            div.appendChild(span3)
            div.appendChild(div1)
            div.appendChild(div2)
            div.appendChild(div3)
            img.src =photo
            h2.innerHTML = result[i].author
            span1.innerHTML = result[i].date
            span2.innerHTML = result[i].article_title
            span3.innerHTML = result[i].article_content
            div1.innerHTML = '点赞' + result[i].like
            div2.innerHTML = '评论' + result[i].comment
            div3.innerHTML = '浏览' + result[i].view

            div.id = result[i].id
            div.onclick = function (e) {
                location.href = '/content?id=' + e.target.id
            }
        } else {
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
            
            img.src =photo
            h2.innerHTML = result[i].author
            span1.innerHTML = result[i].date
            span2.innerHTML = result[i].article_title
            span3.innerHTML = result[i].article_content
            div1.innerHTML = '点赞' + result[i].like

            let ii =result[i].id
            let l=result[i].like
           div1.onclick=async function(){

               let obj = {
                   id: ii
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
                   l++
                   div1.innerHTML = '点赞' + l

               } else {
                   l--
                   div1.innerHTML = '点赞' + l
               }
           }
        }
    }
}
ask_for_data()