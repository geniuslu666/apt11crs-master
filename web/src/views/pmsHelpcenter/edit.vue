<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" :z-index="99">
      <n-drawer-content closable :header-style="{
                    padding: '20px',
                  }" :body-content-style="{
                    padding: '20px',
                  }" :footer-style="{
                    padding: '12px 20px',
                  }">
        <template #header>
          <div style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">{{ formValue.id > 0 ? '编辑帮助中心 #' + formValue.id : '添加帮助中心' }}</div>
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
          <n-form
            ref="formRef"
            label-placement="top"
            label-width="auto"
            :model="formValue"
            :rules="rules"
          >
            <n-grid :cols="1">
              <n-gi>
                <n-form-item label="语言" path="language">
                  <n-select placeholder="请选择语言" v-model:value="formValue.language" :options="options.language" @update:value="handleUpdateLanguageValue" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="分类" path="categoryId">
                  <n-select
                    v-model:value="formValue.categoryId"
                    :options="options.category"
                    clearable
                    filterable
                    label-field="name"
                    value-field="id"
                  />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="标题" path="title">
                  <n-input v-model:value="formValue.title" placeholder="请输入标题" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="内容" path="content">
                  <Editor id="content" v-model:modelValue="formValue.content" style="height: 450px" />
                </n-form-item>
              </n-gi>
              <n-gi>
                <n-form-item label="排序" path="sort">
                  <n-input-number
                    v-model:value="formValue.sort"
                    clearable
                    placeholder="请输入排序"
                    style="width: 100%"
                  />
                  <template #feedback>
                    <div style="font-size: 12px">越大越靠前</div>
                  </template>
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>

      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
import {computed, ref} from 'vue';
import {Edit, View} from '@/api/pmsHelpcenter';
import {newState, options, rules, State} from './model';
import Editor from '@/components/Editor/editor.vue';
import {useMessage} from 'naive-ui';
import {adaModalWidth} from '@/utils/hotgo';
import {All} from "@/api/pmsHelpcenterCategory/index";

const emit = defineEmits(['reloadTable']);
const message = useMessage();
const loading = ref(false);
const showModal = ref(false);
const formValue = ref<State>(newState(null));
const formRef = ref<any>({});
const formBtnLoading = ref(false);
const dialogWidth = computed(() => {
  return adaModalWidth(650);
});

function openModal(state: State) {
  showModal.value = true;

  // 新增
  if (!state || state.id < 1) {
    formValue.value = newState(state);

    return;
  }

  // 编辑
  loading.value = true;

  All({
    language: state.language,
  }).then((res) => {
    options.value.category = res.list;
  });

  View({ id: state.id })
    .then((res) => {
      formValue.value = res;
    })
    .finally(() => {
      loading.value = false;
    });
}

function handleUpdateLanguageValue(value) {
  All({
    language: value,
  }).then((res) => {
    formValue.value.categoryId = null
    options.value.category = res.list;
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

<style lang="less"></style>


