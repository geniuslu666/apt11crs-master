<template>
  <editor v-model="editorValue" :init="init"
          api-key="7a8g2fnpj9juq570cyued3oei9ay6yc766tloiu91qb68pvo"></editor>
</template>

<script setup>
import {reactive, ref, toRefs, watch} from 'vue';
import {UploadImage} from '@/api/base';
import {isNullOrUnDef} from '@/utils/is';
import tinymce from 'tinymce/tinymce.js';
import 'tinymce/models/dom'; // (TinyMCE 6)
// 外观
import 'tinymce/skins/ui/oxide/skin.css';
import 'tinymce/themes/silver';

// Icon
import 'tinymce/icons/default';

// 导入免费版插件
import 'tinymce/plugins/advlist';        // 高级列表
import 'tinymce/plugins/anchor';         // 锚点
import 'tinymce/plugins/autolink';       // 自动链接
import 'tinymce/plugins/autoresize';     // 自动调整大小
import 'tinymce/plugins/autosave';       // 自动保存
import 'tinymce/plugins/charmap';        // 特殊字符
import 'tinymce/plugins/code';           // 源代码
import 'tinymce/plugins/codesample';     // 代码示例
import 'tinymce/plugins/directionality'; // 文字方向
import 'tinymce/plugins/emoticons';      // 表情
import 'tinymce/plugins/emoticons/js/emojis.js';
import 'tinymce/plugins/fullscreen';     // 全屏
import 'tinymce/plugins/help';           // 帮助
import 'tinymce/plugins/image';          // 图片
import 'tinymce/plugins/importcss';      // 导入CSS
import 'tinymce/plugins/insertdatetime'; // 插入日期时间
import 'tinymce/plugins/link';           // 链接
import 'tinymce/plugins/lists';          // 列表
import 'tinymce/plugins/media';          // 媒体
import 'tinymce/plugins/nonbreaking';    // 不间断空格
import 'tinymce/plugins/pagebreak';      // 分页符
import 'tinymce/plugins/preview';        // 预览
import 'tinymce/plugins/quickbars';      // 快速工具栏
import 'tinymce/plugins/save';           // 保存
import 'tinymce/plugins/searchreplace';  // 查找替换
import 'tinymce/plugins/table';          // 表格
import 'tinymce/plugins/visualblocks';   // 可视化块
import 'tinymce/plugins/visualchars';    // 可视化字符
import 'tinymce/plugins/wordcount';      // 字数统计
import 'tinymce/plugins/help/js/i18n/keynav/zh_CN.js';

// 语言包
import 'tinymce-i18n/langs7/zh_CN.js';

// TinyMCE-Vue
import Editor from '@tinymce/tinymce-vue';

console.log(tinymce)

const props = defineProps({
  modelValue: {
    type: String,
    default: '',
  },
  plugins: {
    type: [String, Array],
    default: 'advlist anchor autolink autoresize autosave charmap code codesample directionality emoticons fullscreen help image importcss insertdatetime link lists media nonbreaking pagebreak preview quickbars save searchreplace table visualblocks visualchars wordcount',
  },
  toolbar: {
    type: [String, Array],
    default: 'undo redo | fontfamily fontsize | bold italic underline strikethrough | forecolor backcolor | alignleft aligncenter alignright alignjustify | outdent indent | bullist numlist | link image media table | emoticons charmap codesample | searchreplace | visualblocks visualchars | code | fullscreen help',
  },
});

const emit = defineEmits(['update:modelValue']);

const init = reactive({
      language: 'zh_CN',
      min_height: 326,
      menubar: true,
    // content_css: false,
      content_style: 'p { margin: 0; }',
      skin: false,
      plugins: props.plugins,
      toolbar: props.toolbar,
      font_size_formats: '12px 14px 16px 18px 20px 22px 24px 28px 32px 36px 48px 56px 72px',
      font_family_formats: '微软雅黑=Microsoft YaHei,Helvetica Neue,PingFang SC,sans-serif;苹果苹方=PingFang SC,Microsoft YaHei,sans-serif;宋体=simsun,serif;仿宋体=FangSong,serif;黑体=SimHei,sans-serif;Arial=arial,helvetica,sans-serif;Arial Black=arial black,avant garde;Book Antiqua=book antiqua,palatino;Comic Sans MS=comic sans ms,sans-serif;Courier New=courier new,courier;Georgia=georgia,palatino;Helvetica=helvetica;Impact=impact,chicago;Symbol=symbol;Tahoma=tahoma,arial,helvetica,sans-serif;Terminal=terminal,monaco;Times New Roman=times new roman,times;Trebuchet MS=trebuchet ms,geneva;Verdana=verdana,geneva;Webdings=webdings;Wingdings=wingdings,zapf dingbats',
      branding: false,
      promotion: false,
      relative_urls: false,
      remove_script_host: false,
      toolbar_mode: 'wrap',
      contextmenu: 'link image table',
      image_advtab: true,
      image_description: true,
      image_dimensions: true,
      image_title: true,
      automatic_uploads: true,
      file_picker_types: 'image',
      setup: function(editor) {
        let showImageDialog = false;

        editor.on('init', function() {
          const originalImageAction = editor.ui.registry.getAll().buttons.image.onAction;

          editor.ui.registry.getAll().buttons.image.onAction = function() {
            if (editor.selection.getNode().nodeName === 'IMG') {
              showImageDialog = true;
              originalImageAction();
              return;
            }

            const input = document.createElement('input');
            input.setAttribute('type', 'file');
            input.setAttribute('accept', 'image/*');

            input.onchange = function() {
              const file = this.files[0];

              const formData = new FormData();
              formData.append('file', file);

              UploadImage(formData)
                .then((res) => {
                  editor.insertContent(`<img src="${res.fileUrl}" alt="${file.name}" style="width: 100%;" />`);
                })
                .catch((err) => {
                  console.error('Error:', err);
                });
            };

            input.click();
          };
        });
      },
      file_picker_callback: function(callback, value, meta) {
        if (meta.filetype === 'image') {
          const input = document.createElement('input');
          input.setAttribute('type', 'file');
          input.setAttribute('accept', 'image/*');

          input.onchange = function() {
            const file = this.files[0];

            const formData = new FormData();
            formData.append('file', file);

            UploadImage(formData)
              .then((res) => {
                callback(res.fileUrl, {
                  alt: file.name,
                  style: 'width: 100%'
                });
              })
              .catch((err) => {
                console.error('Error:', err);
              });
          };

          input.click();
        }
      },
    }
  )
;

const {modelValue} = toRefs(props);
const editorValue = ref(modelValue.value);

watch(modelValue, (newValue) => {
  editorValue.value = newValue;
});

watch(editorValue, (newValue) => {
  emit('update:modelValue', newValue);
});

function checkFileType(map, fileType) {
  if (isNullOrUnDef(map)) {
    return true;
  }
  return map.includes(fileType);
}

</script>

<style scoped>

</style>
