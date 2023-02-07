let changing0 = document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[0]
let changing1 = document.querySelector('.r-side').querySelectorAll('div')[0].querySelectorAll('span')[1]
changing0.onclick=function(){
    location.href='/creator/data/follow/data'
}
changing1.onclick=function(){
    location.href='/creator/data/follower/list'
}


let follower_list_detail = document.querySelector('.follower_list_detail')

async function ask_for_follower_list() {
    let obj = {
        user_id: localStorage.getItem('myself_id')
    }
    let res = await fetch('http://localhost:8080/creator/data/follower/list', {
        method: 'post',
        headers: {
            'Content-type': 'application/json'
        },

    });
    let res2 = await res.json();
    let result = res2
    for (let i = 0; i < result.length; i++) {
        let div = document.querySelector('div')
        follower_list_detail.appendChild(div)
        let img = document.querySelector('img')
        let h1 = document.querySelector('h1')
        let button = document.querySelector('button')
        div.appendChild(img)
        div.appendChild(h1)
        div.appendChild(button)
        img.src = result[i].cover
        h1.innerHTML = result[i].name

        let obj2 = {
            phoneoremail: result[i].phoneoremail
        }
        let r = await fetch('http://localhost:8080/getuserid', {
            method: 'post',
            headers: {
                'Content-type': 'application/json'
            },
            body: JSON.stringify(obj2)
        });
        let r2 = await r.json();
        let u_id = r2.msg.user_id

        div.id = u_id
        div.onclick = function (e) {
            location.href = '/web/winter-work/personal information-other/person-home-other/person3.html?id=' + e.target.id
        }


        button.onclick = async function () {
            let obj = {
                followed_id: u_id
            }
            let res = await fetch('http://localhost:8080/notice', {
                method: 'post',
                headers: {
                    'Content-type': 'application/json'
                },
                body: JSON.stringify(obj)
            });
            let res2 = await res.json();
            let result = res2.msg
            if (result == '关注成功') {
                button.innerHTML = '已关注'
            } else if (result == '取消关注') {
                button.innerHTML = '关注'
            }
        }
    }
}
ask_for_follower_list()