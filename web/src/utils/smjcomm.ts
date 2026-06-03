import transjson from '@/utils/lang/translations.json';

import { useUserStore } from '@/store/modules/user';
const userStore = useUserStore();

export function alertSound(soundUrl) {
  var b = document.getElementById('chatMessageAudio');
  if (b.canPlayType('audio/ogg; codecs="vorbis"')) {
    b.type = 'audio/mpeg';
    b.src = soundUrl;
    var p = b.play();
    p && p.then(function () {}).catch(function (e) {});
  }
}

//静态页面 多语言显示
export function translang(val) {
  // const transdata = transjson[val][userStore.language]
  if (userStore.language != 'zh') {
    return transjson[val] ? transjson[val][userStore.language] : val;
  } else {
    return val;
  }
}
// 获取语言信息
export function getlang(langdata, type) {
  let lang = {};
  langdata.forEach((f) => {
    if (f.language == type) {
      lang = f;
    }
  });

  return lang;
}
// 拆分 imgurl 名称
export function imgoutlist(data) {
  let aa = data.map((m) => {
    const name = m.split('/');
    return {
      name: name[name.length - 1],
      url: m,
    };
  });

  return aa;
}
// 转换语言键值对
export function jsontoobj(data) {
  let objlang = {};
  if (data) {
    data.forEach((element) => {
      if (element.content.indexOf('|') != -1) {
        element.listok = true;
        element.content = element.content.split('|');
      } else {
        element.listok = false;
      }
      objlang[element.language] = element;
    });

    return objlang;
  }
  return {};
}
