<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="'发送站内信'"
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
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="标题" path="value">
                  <n-input v-model:value="formValue.title" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="内容" path="value">
                  <n-input v-model:value="formValue.content" type="textarea" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
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
  import { SendMsg } from '@/api/pmsMember';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const settingStore = useProjectSettingStore();
  const loading = ref(false);
  const showModal = ref(false);
  const formValue = reactive({
    content: '',
    member_id: '',
    title: '',
  });
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(650);
  });

  function openModal(data) {
    showModal.value = true;
    formValue.content = '';
    formValue.member_id = data;
  }

  function confirmForm(e) {
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        SendMsg(formValue).then((_res) => {
          message.success('发送成功');
          setTimeout(() => {
            closeForm();
            emit('reloadTable');
          });
        });
      } else {
        message.error('请填写完整信息');
      }
      formBtnLoading.value = false;
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
