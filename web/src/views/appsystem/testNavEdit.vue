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
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑导航 #' + formValue.id : '添加导航' }}</div>
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
              <n-grid :cols="1">
                <n-gi>
                  <n-form-item label="导航名称" path="name" :show-require-mark="true">
                    <n-input placeholder="请输入导航名称" v-model:value="formValue.name" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>

                <n-gi>
                  <n-form-item label="图标" path="image" :show-feedback='false'>
                    <UploadImage :maxNumber="1" v-model:value="formValue.image" />
                  </n-form-item>
                  <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px">建议尺寸：48px*48px</div>
                </n-gi>
                <n-gi>
                  <n-form-item label="App链接" path="appLink" :show-require-mark="true">
                    <n-input placeholder="请输入App链接" v-model:value="formValue.appLink" :style="{ width: '300px' }" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="小程序链接" path="wxLink" :show-require-mark="true">
                    <n-input placeholder="请输入小程序链接" v-model:value="formValue.wxLink" :style="{ width: '300px' }" />
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
import { Edit, View } from '@/api/pmsTestNav';
import { State, newState } from './test_nav_model';
import { useMessage } from 'naive-ui';
import {adaModalWidth} from "@/utils/hotgo";

const emit = defineEmits(['reloadTable']);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});
const message = useMessage();
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


