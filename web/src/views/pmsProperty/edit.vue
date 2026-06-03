<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth">
      <n-drawer-content
        :title="formValue.id > 0 ? '编辑物业 #' + formValue.id : '添加物业'"
        closable
      >
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            :label-placement="settingStore.isMobile ? 'top' : 'left'"
            :label-width="100"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:2 l:2 xl:2 2xl:2" responsive="screen">
              <n-gi span="1">
                <n-form-item label="图标" path="icon">
                  <FileChooser :maxNumber="1" v-model:value="formValue.icon" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="封面" path="cover">
                  <FileChooser :maxNumber="1" v-model:value="formValue.cover" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="货币" path="currency">
                  <n-input placeholder="请输入货币" v-model:value="formValue.currency" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="语言" path="language">
                  <n-input placeholder="请输入语言" v-model:value="formValue.language" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="时区" path="timeZone">
                  <n-input placeholder="请输入时区" v-model:value="formValue.timeZone" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="联系人" path="contactName">
                  <n-input placeholder="请输入联系人" v-model:value="formValue.contactName" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="联系方式" path="phone">
                  <n-input placeholder="请输入联系方式" v-model:value="formValue.phone" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="邮箱" path="contactEmail">
                  <n-input placeholder="请输入邮箱" v-model:value="formValue.contactEmail" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最大预定区间" path="maxDaysNotice">
                  <n-input-number
                    placeholder="请输入最大预定区间"
                    v-model:value="formValue.maxDaysNotice"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="退房后 分钟" path="minutesAfterCheckout">
                  <n-input-number
                    placeholder="请输入退房后 分钟"
                    v-model:value="formValue.minutesAfterCheckout"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="入住前 分钟" path="minutesBeforeCheckin">
                  <n-input-number
                    placeholder="请输入入住前 分钟"
                    v-model:value="formValue.minutesBeforeCheckin"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="预约期限" path="bookingLeadTimeLabel">
                  <n-input
                    placeholder="请输入预约期限"
                    v-model:value="formValue.bookingLeadTimeLabel"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="周转天数" path="turnoverDays">
                  <n-input-number
                    placeholder="请输入周转天数"
                    v-model:value="formValue.turnoverDays"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="入住时间" path="checkinAt">
                  <n-input placeholder="请输入入住时间" v-model:value="formValue.checkinAt" />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="退房时间" path="checkoutAt">
                  <n-input placeholder="请输入退房时间" v-model:value="formValue.checkoutAt" />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>

        <div class="text-r">
          <n-button class="mr-2" @click="closeForm"> 取消 </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm"> 确定 </n-button>
        </div>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { ref, computed } from 'vue';
  import { Edit, View } from '@/api/pmsProperty';
  import { State, newState, rules } from './model';
  import { useProjectSettingStore } from '@/store/modules/projectSetting';
  import { useMessage } from 'naive-ui';
  import { adaModalWidth } from '@/utils/hotgo';

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
