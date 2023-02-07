let column_head = document.querySelector('.column_head')
let hd_h1 = column_head.querySelector('h1')
let hd_img = column_head.querySelector('img')
let hd_span = column_head.querySelector('span')
let hd_div = column_head.querySelector('div')
let my_id = new URL(location.href).searchParams.get("id")
let article_list = document.querySelector('.article_list')
async function ask_for_detail() {
    let obj = {
        column_id: my_id
    }
    let res = await fetch('http://localhost:8080/column', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },
        body: JSON.stringify(obj)
    })
    let res2 = await res.json();
    let result1 = res2.column;
    hd_h1.innerHTML = result1.column_title;
    hd_img.src = result1.cover;
    hd_span.innerHTML = result1.phoneoremail;
    hd_div.innerHTML = result1.column_intruduction;
    let result2 = res2.articles;

    for (var i = 0; i < result2.length; i++) {
        let div = document.createElement('div');
        let h2 = document.createElement('h2');
        let span1 = document.createElement('span');
        let span2 = document.createElement('span');
        let span3 = document.createElement('span');
        let span4 = document.createElement('span');
        let span5 = document.createElement('span');
        let span6 = document.createElement('span');
        article_list.appendChild(div)
        div.appendChild(span1)
        div.appendChild(span2)
        div.appendChild(h2)
        div.appendChild(span3)
        div.appendChild(span4)
        div.appendChild(span5)
        div.appendChild(span6)
        let a=result2[i].article_title
        let b=result2[i].article_content
        span1.innerHTML = result2[i].author
        span2.innerHTML = result2[i].category
        h2.innerHTML = a
        span3.innerHTML = b
        span4.innerHTML ='喜欢'+ result2[i].like
        span5.innerHTML ='浏览'+ result2[i].view
        span6.innerHTML = '收藏'+result2[i].collection
        div.id=result2[i].id

        span1.id=result2[i].id;
        span2.id=result2[i].id;
        span3.id=result2[i].id;

        div.onclick=function(e){
            location.href='/content?id='+e.target.id
        }
        h2.id=result2[i].id
        h2.onclick=function(e){
            location.href='/content?id='+e.target.id
        }

    }

}
ask_for_detail()