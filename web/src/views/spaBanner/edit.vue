<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑Banner #' + formValue.id : '添加Banner'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="语言" path="language">
                  <n-radio-group v-model:value="formValue.language" name="language">
                    <n-radio-button
                      v-for="language in options.language"
                      :key="language.value"
                      :value="language.value"
                      :label="language.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="轮播图" path="bannerImage" :show-feedback='false'>
                  <FileChooser :maxNumber="1" v-model:value="formValue.bannerImage" />
                </n-form-item>
                <div style="color:#9EA4AA; font-size: 12px;line-height: 34px;margin: 0 0 24px;padding-left: 100px;">建议尺寸：750px*826px</div>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="排序" path="sort">
                  <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" style="width: 100px"/>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button
                      v-for="type in options.sys_normal_disable"
                      :key="type.value"
                      :value="type.value"
                      :label="type.label"
                    />
                  </n-radio-group>
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
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm">
            确定
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { Edit, View, MaxSort } from '@/api/spaBanner';
import { options, State, newState } from './model';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';

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

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    loading.value = true;
    MaxSort()
      .then((res) => {
        formValue.value.sort = res.sort;
      })
      .finally(() => {
        loading.value = false;
      });
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


