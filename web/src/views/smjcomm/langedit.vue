<template>
  <div v-if="showModal">
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="title"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar v-if="langtype != 'TagList'" style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <!-- 暂时没有数据源可选择，前端标签显示 只能用特殊符号做分割 -->
          <!-- <n-alert title=" 标签格式可用 ' | ' 符号分割编写" type="info" :bordered="false">
          </n-alert> -->
          <n-tabs type="line" :default-value="tabstype" animated>
            <n-tab-pane tab="日本语" @click="tabclick('ja')" name="ja">
              <n-input
                v-if="isEditor == 'input'"
                placeholder="日本语"
                v-model:value="formValue.ja"
                type="textarea"
              />

              <Editor
                v-else-if="isEditor == 'editor'"
                placeholder="日本语"
                style="height: 300px"
                v-model:modelValue="formValue.ja"
              />
            </n-tab-pane>
            <n-tab-pane tab="English" @click="tabclick('en')" name="en">
              <n-input
                v-if="isEditor == 'input'"
                placeholder="English"
                v-model:value="formValue.en"
                type="textarea"
              />
              <Editor
                v-else-if="isEditor == 'editor'"
                placeholder="English"
                style="height: 300px"
                v-model:modelValue="formValue.en"
              />
            </n-tab-pane>
            <n-tab-pane tab="简体中文" @click="tabclick('zh')" name="zh">
              <n-input
                v-if="isEditor == 'input'"
                placeholder="简体中文"
                v-model:value="formValue.zh"
                type="textarea"
              />
              <Editor
                v-else-if="isEditor == 'editor'"
                placeholder="简体中文"
                style="height: 300px"
                v-model:modelValue="formValue.zh"
              />
            </n-tab-pane>
            <n-tab-pane tab="한국어" @click="tabclick('ko')" name="ko">
              <n-input
                v-if="isEditor == 'input'"
                placeholder="한국어"
                v-model:value="formValue.ko"
                type="textarea"
              />
              <Editor
                v-else-if="isEditor == 'editor'"
                placeholder="한국어"
                style="height: 300px"
                v-model:modelValue="formValue.ko"
              />
            </n-tab-pane>
            <n-tab-pane tab="繁体中文" @click="tabclick('zh_CN')" name="zh_CN">
              <n-input
                v-if="isEditor == 'input'"
                placeholder="繁体中文"
                v-model:value="formValue.zh_CN"
                type="textarea"
              />
              <Editor
                v-else-if="isEditor == 'editor'"
                placeholder="繁体中文"
                style="height: 300px"
                v-model:modelValue="formValue.zh_CN"
              />
            </n-tab-pane>
          </n-tabs>
        </n-spin>
      </n-scrollbar>
      <n-scrollbar v-else style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-tabs type="line" :default-value="tabstype" animated>
            <n-tab-pane tab="日本语" @click="tabclick('ja')" name="ja">
              <n-dynamic-tags v-model:value="formValuetag.ja" />
            </n-tab-pane>
            <n-tab-pane tab="English" @click="tabclick('en')" name="en">
              <n-dynamic-tags v-model:value="formValuetag.en" />
            </n-tab-pane>
            <n-tab-pane tab="简体中文" @click="tabclick('zh')" name="zh">
              <n-dynamic-tags v-model:value="formValuetag.zh" />
            </n-tab-pane>
            <n-tab-pane tab="한국어" @click="tabclick('ko')" name="ko">
              <n-dynamic-tags v-model:value="formValuetag.ko" />
            </n-tab-pane>
            
            <n-tab-pane tab="繁体中文" @click="tabclick('zh_CN')" name="zh_CN">
              <n-dynamic-tags v-model:value="formValuetag.zh_CN" />
            </n-tab-pane>
          </n-tabs>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm"> 取消 </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm"> 确定 </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed, reactive } from 'vue';
  import { Edit, View, editPropertyLanguage } from '@/api/pmsProperty';
  import UploadImage from '@/components/Upload/uploadImage.vue';
  import UploadFile from '@/components/Upload/uploadFile.vue';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';
  import Editor from '@/components/Editor/editor.vue';

  import { State } from '../curdDemo/model';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const alldata = ref({});
  let title = ref('');
  let langtype = ref('');
  let isEditor = ref('input');
  let tabstype = ref('zh');
  let formValue = reactive({
    en: '',
    zh: '',
    ja: '',
    ko: '',
    zh_CN: '',
  });
  let formValuetag = reactive({
    en: [],
    zh: [],
    ja: [],
    ko: [],
    zh_CN: [],
  });
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(840);
  });

  function openModal(row, type, name, datajson, editor) {
    showModal.value = true;
    langtype = type;
    isEditor = editor ? editor : 'input';
    title = name;
    if (langtype == 'TagList') {
      formValuetag.en = datajson.en ? JSON.parse(datajson.en.content) : [];
      formValuetag.zh = datajson.zh ? JSON.parse(datajson.zh.content) : [];
      formValuetag.ja = datajson.ja ? JSON.parse(datajson.ja.content) : [];
      formValuetag.ko = datajson.ko ? JSON.parse(datajson.ko.content) : [];
      formValuetag.zh_CN = datajson.zh_CN ? JSON.parse(datajson.zh_CN.content) : [];
    } else {
      formValue.en = datajson.en ? datajson.en.content : '';
      formValue.zh = datajson.zh ? datajson.zh.content : '';
      formValue.ja = datajson.ja ? datajson.ja.content : '';
      formValue.ko = datajson.ko ? datajson.ko.content : '';
      formValue.zh_CN = datajson.zh_CN ? datajson.zh_CN.content : '';
    }
    // formValue.zh_CN = datajson.zh_CN.content;

    // 新增
    if (!row || row.id < 1) {
      alldata.value = row;

      return;
    }

    alldata.value = row;

    // 编辑
    // loading.value = true;
    // View({ id: row.id })
    //   .then((res) => {
    //     alldata.value = res;
    //   })
    //   .finally(() => {
    //     loading.value = false;
    //   });
  }
  function tabclick(type) {
    tabstype = type;
  }
  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;

    let dataobj = {
      type: langtype,
      id: alldata.value.id,
      Data: {
        name: {},
        description: {},
        address: {},
        tag_list: {},
        room_des: {},
        surroundings: {},
        bus_station: {},
        subway: {},
        required_book: {},
        cancel_policy: {},
      },
    };
    if (langtype == 'Name') {
      dataobj.Data.name = formValue;
    } else if (langtype == 'Description') {
      dataobj.Data.description = formValue;
    } else if (langtype == 'Address') {
      dataobj.Data.address = formValue;
    } else if (langtype == 'TagList') {
      const aa = JSON.parse(JSON.stringify(formValuetag));
      aa.en = JSON.stringify(aa.en);
      aa.zh = JSON.stringify(aa.zh);
      aa.ja = JSON.stringify(aa.ja);
      aa.ko = JSON.stringify(aa.ko);
      aa.zh_CN = JSON.stringify(aa.zh_CN);
      dataobj.Data.tag_list = aa;
    } else if (langtype == 'RoomDes') {
      dataobj.Data.room_des = formValue;
    } else if (langtype == 'Surroundings') {
      dataobj.Data.surroundings = formValue;
    } else if (langtype == 'BusStation') {
      dataobj.Data.bus_station = formValue;
    } else if (langtype == 'Subway') {
      dataobj.Data.subway = formValue;
    } else if (langtype == 'RequiredBook') {
      dataobj.Data.required_book = formValue;
    } else if (langtype == 'CancelPolicy') {
      dataobj.Data.cancel_policy = formValue;
    }
    editPropertyLanguage(dataobj).then((_res) => {
      message.success('操作成功');
      setTimeout(() => {
        closeForm();
        emit('reloadTable');
      });
    });

    formBtnLoading.value = false;
  }

  function closeForm() {
    showModal.value = false;
    loading.value = false;
  }

  defineExpose({
    openModal,
  });
</script>

<style lang="less"></style>
