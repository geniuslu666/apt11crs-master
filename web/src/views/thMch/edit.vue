<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑商户 #' + formValue.id : '添加商户' }}</div>
        </template>
        <template #footer>
          <n-button @click="closeForm" style="width: 70px;height: 35px;margin-right: 10px">
            取消
          </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm" style="width: 70px;height: 35px;">
            保存
          </n-button>
        </template>
        <n-spin :show="loading" description="请稍候...">
          <n-form :model="formValue" ref="formRef"  label-placement="top"
                  label-width="auto"
                  require-mark-placement="right-hanging"
          >
            <div class="level-detail-div">
              <div class="level-detail-div-title">
                <div></div>
                基础信息
              </div>
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="商户名称" path="name">
                    <n-input placeholder="请输入商户名称" v-model:value="formValue.name" style="width: 300px" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="商户分类" path="categoryId">
                    <n-select
                      v-model:value="formValue.categoryId"
                      :options="cateList"
                      clearable
                      filterable
                      label-field="name"
                      value-field="id"
                      style="width: 300px"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="商户LOGO" path="logo" :show-feedback='false'>
                    <UploadImage :maxNumber="1" v-model:value="formValue.logo" />
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px">建议尺寸：140px*140px</div>
                </n-gi>
                <n-gi>
                  <n-form-item label="联系信息" path="contactInfo">
                    <n-input type="textarea" placeholder="请输入联系信息" v-model:value="formValue.contactInfo" style="width: 300px" />
                  </n-form-item>
                </n-gi>
              </n-grid>
            </div>
          </n-form>
        </n-spin>

      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Edit, View } from '@/api/thMch';
import { State, newState, cateList } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import UploadImage from "@/components/Upload/uploadImage.vue";
import {adaModalWidth} from "@/utils/hotgo";

const emit = defineEmits(['reloadTable']);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const message = useMessage();
const settingStore = useProjectSettingStore();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);

    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      if(!formValue.value.name){
        message.error('请填写商户名称');
        formBtnLoading.value = false;
        return false;
      }
      if(!formValue.value.categoryId){
        message.error('请选择商户分类');
        formBtnLoading.value = false;
        return false;
      }
      if(!formValue.value.logo){
        message.error('请上传LOGO');
        formBtnLoading.value = false;
        return false;
      }
      // if(!formValue.value.contactInfo){
      //   message.error('请填写联系信息');
      //   formBtnLoading.value = false;
      //   return false;
      // }
      Edit(formValue.value).then((_res) => {
        message.success('操作成功');
        setTimeout(() => {
          formBtnLoading.value = false;
          closeForm();
          emit('reloadTable');
        });
      }).catch((err) => {
        formBtnLoading.value = false;
      });
    } else {
      message.error('请填写完整信息');
      formBtnLoading.value = false;
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

<style lang="less">
.level-detail-div{
  &-title{
    display: flex;
    align-items: center;
    font-weight: 500;
    font-size: 16px;
    color: #3D3D3D;
    line-height: 22px;
    margin-bottom: 15px;
    div{
      margin-right: 5px;
      width: 6px;
      height: 15px;
      background: #053DC8;
    }
  }
  &-item{
    &-title{
      font-weight: 400;
      font-size: 14px;
      color: #3D3D3D;
      line-height: 20px;
      margin-bottom: 8px;
    }
    &-content{

    }
  }
}
</style>


