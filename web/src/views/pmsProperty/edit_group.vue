<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      title="绑定会员分组"
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
                <n-form-item label="会员分组" path="group">
                  <n-select
                    placeholder="请选择会员分组"
                    v-model:value="formValue.groupIds"
                    :options="groupList"
                    label-field="memberGroup"
                    value-field="id"
                    clearable
                    filterable
                    multiple
                  />
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
import {computed, ref} from 'vue';
import {GroupAll} from '@/api/pmsMemberGroup';
import {EditGroup} from '@/api/pmsProperty';
import {newState,  State} from './model';
import {useProjectSettingStore} from '@/store/modules/projectSetting';
import {useMessage} from 'naive-ui';
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

const groupList = ref([])

function openModal(state: State) {
  showModal.value = true;
  formValue.value = newState(state);
  if(state.groupIds != '') {
    formValue.value.groupIds = state.groupIds.split(",").map(item => Number(item));
  }
  // console.log('formValue11', formValue.value);

  loading.value = true;
  GroupAll({}).then((res) => {
    groupList.value = res.list;
    loading.value = false;
  })
  .finally(() => {
    loading.value = false;
  });

  // 编辑
  // loading.value = true;
  // View({ id: state.id })
  //   .then((res) => {
  //     console.log('groupIds', res.groupIds);
  //     if(res.groupIds != ''){
  //       res.groupIds = JSON.parse(res.groupIds);
  //     }else{
  //       res.groupIds = [];
  //     }
  //
  //     formValue.value = res;
  //     console.log('formValue', formValue.value);
  //   })
  //   .finally(() => {
  //     loading.value = false;
  //   });
}

function confirmForm(e) {
  e.preventDefault();
  formBtnLoading.value = true;
  formRef.value.validate((errors) => {
    if (!errors) {
      EditGroup(formValue.value).then((_res) => {
        message.success('操作成功');
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


