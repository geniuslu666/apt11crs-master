<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑商家通知 #' + formValue.id : '添加商家通知'"
      :style="{
        width: dialogWidth,
      }"
    >
      <n-scrollbar style="max-height: 87vh" class="pr-5">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="150"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="通知标题" path="title">
                  <n-input placeholder="请输入通知标题" v-model:value="formValue.title" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="公告内容" path="content">
                  <Editor style="height: 300px" id="content" v-model:modelValue="formValue.content" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="排序" path="sort">
                  <n-input-number placeholder="请输入排序" v-model:value="formValue.sort" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="状态" path="status">
                  <n-radio-group v-model:value="formValue.status" name="status">
                    <n-radio-button
                      v-for="status in options.sys_normal_disable"
                      :key="status.value"
                      :value="status.value"
                      :label="status.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="是否需要用户确认" path="needUserConfirm">
                  <n-switch :unchecked-value="2" :checked-value="1" v-model:value="formValue.needUserConfirm"/>
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
import { Edit, View, MaxSort } from '@/api/foodRestaurantNotice';
import { options, State, newState, rules } from './model';
import Editor from '@/components/Editor/editor.vue';
import { useProjectSettingStore } from '@/store/modules/projectSetting';
import { useMessage } from 'naive-ui';
import { adaModalWidth } from '@/utils/hotgo';

const props = defineProps({
  restaurantId: {
    type: Number,
    default: 0,
  },
});
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
  loading.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);
    formValue.value.restaurantId = parseInt(props.restaurantId)

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


