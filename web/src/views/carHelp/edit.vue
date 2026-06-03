<template>
  <div>
    <a-modal
      v-model:open="showModal"
      :mask-closable="false"
      :show-icon="false"
      :style="{
        width: dialogWidth,
      }"
      :title="formValue.id > 0 ? '编辑帮助说明 #' + formValue.id : '添加帮助说明'"
      :z-index="99"
      preset="card"
      transform-origin="center"
      @ok="confirmForm"
    >
      <n-scrollbar class="pr-5" style="max-height: 87vh">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            :model="formValue"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="语言" path="language">
                  <n-select
                    v-model:value="formValue.language"
                    :options="languageOptions"
                    style="width: 150px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="主图" path="image" :show-feedback='false'>
                  <FileChooser v-model:value="formValue.image" :maxNumber="1" />
                </n-form-item>
                <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 100px;">建议尺寸：702px*174px</div>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="内容" path="content">
                  <Editor id="content" v-model:modelValue="formValue.content" style="height: 450px" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
      </n-scrollbar>
      <template #action>
        <n-space>
          <n-button @click="closeForm">
            取消
          </n-button>
          <n-button :loading="formBtnLoading" type="info" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </a-modal>
  </div>
</template>

<script lang="ts" setup>
import {computed, reactive, ref} from 'vue';
import {Edit, LanguageList, View} from '@/api/carHelp';
import {newState, State} from './model';
import {useProjectSettingStore} from '@/store/modules/projectSetting';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import Editor from "@/components/Editor/editor.vue";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(840);
});
const languageOptions = reactive([
  {
    value: 'zh',
    label: '中文',
    disabled: false,
  },
  {
    value: 'en',
    label: '英语',
    disabled: false,
  },
  {
    value: 'ja',
    label: '日文',
    disabled: false,
  },
  {
    value: 'ko',
    label: '韩文',
    disabled: false,
  },
  {
    value: 'zh_CN',
    label: '繁体',
    disabled: false,
  }
]);

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    loading.value = true;
    LanguageList({})
      .then((res) => {
        let initLanguageList = res.list.map((item)=>{
          return item.language
        })
        languageOptions.forEach((item)=>{
          if(initLanguageList.indexOf(item.value) >= 0){
            item.disabled = true
          }else{
            item.disabled = false
          }
        })
        loading.value = false;
      })
    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
      LanguageList({})
        .then((res1) => {
          let initLanguageList = res1.list.map((item)=>{
            return item.language
          })
          languageOptions.forEach((item)=>{
            if(initLanguageList.indexOf(item.value) >= 0 && item.value != res.language){
              item.disabled = true
            }else{
              item.disabled = false
            }
          })
          loading.value = false;
        })
    })
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!formValue.value.language){
        formBtnLoading.value = false;
        message.error('请选择语言');
        return false;
      }
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        formBtnLoading.value = false;
        setTimeout(() => {
          closeForm();
          emit('reloadTable');
        });
      }).catch((err)=>{
        formBtnLoading.value = false;
      });
    } else {
      formBtnLoading.value = false;
      message.error('请填写完整信息');
    }
  });
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


