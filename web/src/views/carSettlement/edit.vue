<template>
  <div>
    <n-modal
      v-model:show="showModal"
      :mask-closable="false"
      :show-icon="false"
      preset="dialog"
      transform-origin="center"
      :title="formValue.id > 0 ? '编辑结算模式 #' + formValue.id : '添加结算模式'"
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
                <n-form-item label="名称" path="name_zh">
                  <n-input placeholder="请输入结算模式名称" v-model:value="formValue.name" style="width: 150px" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="结算方式 " path="type">
                  <n-radio-group v-model:value="formValue.type" name="type">
                    <n-radio-button
                      v-for="type in options.settlement_type"
                      :key="type.value"
                      :value="type.value"
                      :label="type.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type == 2">
                <n-form-item label="结算周期" path="cycle">
                  <n-radio-group v-model:value="formValue.cycle" name="cycle">
                    <n-radio-button
                      v-for="cycle in options.settlement_citcle"
                      :key="cycle.value"
                      :value="cycle.value"
                      :label="cycle.label"
                    />
                  </n-radio-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type != 1">
                <n-form-item label="结算比例" path="rate">
                  <n-input-group>
                    <n-input-number placeholder="请输入" :min="0" :max="100" v-model:value="formValue.rate" :show-button="false" style="width: 80px"/>
                    <n-input-group-label>%</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.type != 1">
                <n-form-item label="结算成本控制" path="costArr">
                  <n-checkbox-group v-model:value="formValue.costArr">
                    <n-space>
                      <n-checkbox
                        v-for="item in options.settlement_cost"
                        :key="item.value"
                        :value="item.value"
                        :label="item.label"
                      />
                    </n-space>
                  </n-checkbox-group>
                </n-form-item>
              </n-gi>
<!--              <n-gi span="1" v-if="formValue.type != 1">-->
<!--                <n-form-item label="是否提现审核" path="isAuditWithdraw">-->
<!--                  <n-switch :unchecked-value="2" :checked-value="1" v-model:value="formValue.isAuditWithdraw"-->
<!--                  />-->
<!--                </n-form-item>-->
<!--              </n-gi>-->
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
import { Edit, View } from '@/api/carSettlement';
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

    return;
  }

  // 编辑
  loading.value = true;
  View({ id: state.id })
    .then((res) => {
      res.costArr = res.cost ? res.cost.split(',').map((item)=>{
        return parseInt(item)
      }) : []
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
      formValue.value.cost = formValue.value.costArr ? formValue.value.costArr.join(',') : ''
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


