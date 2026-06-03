<template>
  <div>
    <n-drawer v-model:show="showModal" :width="dialogWidth" @mask-click="closeForm">
      <n-drawer-content :title="translang('预定设置')" closable :native-scrollbar="false">
        <n-spin :show="loading" description="请稍候...">
          <n-form
            ref="formRef"
            :model="formValue"
            :rules="rules"
            label-placement="top"
            label-width="auto"
            class="py-4"
          >
            <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
              <n-gi span="1">
                <n-form-item label="货币" path="currency">
                  <n-input
                    placeholder="请输入货币"
                    v-model:value="formValue.currency"
                    style="width: 120px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="语言" path="language">
                  <n-select
                    placeholder="请选择语言"
                    v-model:value="formValue.language"
                    :options="languageList"
                    style="width: 120px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="时区" path="timeZone">
                  <n-input
                    placeholder="请输入时区"
                    v-model:value="formValue.timeZone"
                    style="width: 120px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最大预定区间" path="maxDaysNotice">
                  <n-input-group>
                    <n-input-number
                      min="0"
                      placeholder="预定区间"
                      v-model:value="formValue.maxDaysNotice"
                      style="width: 120px"
                    />
                    <n-input-group-label>天</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="最小预定区间" path="maxDaysNotice">
                  <n-input-group>
                    <n-input-number
                      min="0"
                      placeholder="预定区间"
                      v-model:value="formValue.minDaysNotice"
                      style="width: 120px"
                    />
                    <n-input-group-label>天</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="联系人" path="contactName">
                  <n-input
                    placeholder="请输入联系人"
                    v-model:value="formValue.contactName"
                    style="width: 220px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="联系方式" path="phone">
                  <n-input
                    placeholder="请输入联系方式"
                    v-model:value="formValue.phone"
                    style="width: 220px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="电子邮箱" path="contactEmail">
                  <n-input
                    placeholder="请输入电子邮箱"
                    v-model:value="formValue.contactEmail"
                    style="width: 220px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="退房后" path="minutesBeforeCheckin">
                  <n-input-group>
                    <n-input-number
                      min="0"
                      v-model:value="formValue.minutesBeforeCheckin"
                      style="width: 120px"
                    />
                    <n-input-group-label>分钟</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="预定期限" path="bookingLeadTimeLabel">
                  <n-input-group>
                    <n-input
                      placeholder="预定期限"
                      v-model:value="formValue.bookingLeadTimeLabel"
                      style="width: 120px"
                    />
                    <n-input-group-label>天</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="周转天数" path="turnoverDays">
                  <n-input-group>
                    <n-input-number
                      min="0"
                      placeholder="周转天数"
                      v-model:value="formValue.turnoverDays"
                      style="width: 120px"
                    />
                    <n-input-group-label>天</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="入住前" path="minutesAfterCheckout">
                  <n-input-group>
                    <n-input-number
                      min="0"
                      v-model:value="formValue.minutesAfterCheckout"
                      style="width: 120px"
                    />
                    <n-input-group-label>分钟</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="入住时间" path="checkinAt">
                  <n-input-group>
                    <n-time-picker
                      placeholder="入住时间"
                      format="HH:mm"
                      v-model:default-formatted-value="formValue.checkinAt"
                      @update:formatted-value="checkinAtChange"
                      @update:value="checkinAtChange1"
                      style="width: 120px"
                    />
                    <n-input-group-label>GMT9+</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="退房时间" path="checkoutAt">
                  <n-input-group>
                    <n-time-picker
                      placeholder="退房时间"
                      format="HH:mm"
                      v-model:default-formatted-value="formValue.checkoutAt"
                      @update:formatted-value="checkoutAtChange"
                      @update:value="checkoutAtChange1"
                      style="width: 120px"
                    />
                    <n-input-group-label>GMT9+</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </n-gi>
              <n-gi span="1">
                <n-form-item label="可见性" path="showType">
                  <n-select
                    placeholder="请选择"
                    v-model:value="formValue.showType"
                    :options="showTypeList"
                    style="width: 220px"
                  />
                </n-form-item>
              </n-gi>
              <n-gi span="1" v-if="formValue.showType == 2">
                <n-form-item label="会员分组" path="groupIdsArr">
                  <n-select
                    placeholder="请选择会员分组"
                    v-model:value="formValue.groupIdsArr"
                    :options="groupList"
                    label-field="memberGroup"
                    value-field="id"
                    clearable
                    filterable
                    multiple
                    style="width: 220px"
                  />
                </n-form-item>
              </n-gi>
            </n-grid>
          </n-form>
        </n-spin>
        <n-space>
          <n-button @click="closeForm"> 取消 </n-button>
          <n-button type="info" :loading="formBtnLoading" @click="confirmForm"> 确定 </n-button>
        </n-space>
      </n-drawer-content>
    </n-drawer>
  </div>
</template>

<script lang="ts" setup>
  import { computed, ref } from 'vue';
  import { Edit } from '@/api/pmsProperty';
  import { FormItemRule, useMessage } from 'naive-ui';
  import { translang } from '@/utils/smjcomm';
  import { adaModalWidth } from '@/utils/hotgo';
  import { Dicts } from '@/api/dict/dict';
  import { GroupAll } from '@/api/pmsMemberGroup';

  const emit = defineEmits(['reloadTable']);
  const message = useMessage();
  const loading = ref(false);
  const showModal = ref(false);
  const dialogWidth = computed(() => {
    return adaModalWidth(620);
  });
  const formValue = ref({
    id: 0,
    maxDaysNotice: 0,
    minDaysNotice: 0,
    turnoverDays: 0,
    checkinAt: '',
    checkoutAt: '',
    checkinAtTime: 0,
    checkoutAtTime: 0,
    currency: '',
    language: '',
    timeZone: '',
    contactName: '',
    phone: '',
    contactEmail: '',
    minutesBeforeCheckin: 0,
    bookingLeadTimeLabel: '',
    minutesAfterCheckout: 0,
    showType: 1,
    groupIdsArr: [],
    groupIds: '',
  });
  const formRef = ref<any>({});
  const formBtnLoading = ref(false);
  const languageList = ref([]);
  const showTypeList = ref([
    {
      label: '全部可见',
      value: 1,
    },
    {
      label: '部分可见',
      value: 2,
    },
  ]);
  const groupList = ref([]);
  const rules = ref({
    groupIdsArr: {
      required: true,
      trigger: ['blur', 'input'],
      validator(rule: FormItemRule, value: string) {
        if (formValue.value.showType == 2 && value.length <= 0) {
          return new Error('请选择会员分组');
        }
        return true;
      },
    },
  });

  function openModal(data) {
    
    loading.value = true; 
    let currentDate = new Date();
    let year = currentDate.getFullYear();
    let month = currentDate.getMonth() + 1;
    let day = currentDate.getDate();
    let formattedDate = `${year}-${month < 10 ? '0' + month : month}-${day < 10 ? '0' + day : day}`;
    data.checkinAt = data.checkinAt ? data.checkinAt : '00:00';
    data.checkoutAt = data.checkoutAt ? data.checkoutAt : '00:00';

    let checkinAtDate = formattedDate + ' ' + data.checkinAt + ':00';
    let checkinAtTime = Date.parse(`${new Date(checkinAtDate)}`);
    let checkoutAtDate = formattedDate + ' ' + data.checkoutAt + ':00';
    let checkoutAtTime = Date.parse(`${new Date(checkoutAtDate)}`);

    Dicts({
      types: ['language'],
    }).then((res) => {
      languageList.value = res.language;
    });

    GroupAll({}).then((res) => {
      groupList.value = res.list;
    });

    formValue.value.id = data.id;
    formValue.value.maxDaysNotice = data.maxDaysNotice;
    formValue.value.minDaysNotice = data.minDaysNotice;
    formValue.value.turnoverDays = data.turnoverDays;
    formValue.value.checkinAt = data.checkinAt;
    formValue.value.checkoutAt = data.checkoutAt;
    formValue.value.checkinAtTime = checkinAtTime;
    formValue.value.checkoutAtTime = checkoutAtTime;
    formValue.value.currency = data.currency;
    formValue.value.language = data.language;
    formValue.value.timeZone = data.timeZone;
    formValue.value.contactName = data.contactName;
    formValue.value.phone = data.phone;
    formValue.value.contactEmail = data.contactEmail;
    formValue.value.minutesBeforeCheckin = data.minutesBeforeCheckin;
    formValue.value.bookingLeadTimeLabel = data.bookingLeadTimeLabel;
    formValue.value.minutesAfterCheckout = data.minutesAfterCheckout;
    let groupIdsArr = data.groupIds ? data.groupIds.split(',').map((item) => Number(item)) : [];
    formValue.value.groupIdsArr = groupIdsArr;
    formValue.value.showType = groupIdsArr.length > 0 ? 2 : 1;
    formValue.value.groupIds = data.groupIds;
    loading.value = false;
    showModal.value = true;
  }

  function checkinAtChange(value) {
    formValue.value.checkinAt = value;
  }

  function checkoutAtChange(value) {
    formValue.value.checkoutAt = value;
  }

  function checkinAtChange1(value) {
    formValue.value.checkinAtTime = value;
  }

  function checkoutAtChange1(value) {
    formValue.value.checkoutAtTime = value;
  }

  function confirmForm(e) {
    if (formValue.value.showType == 1) {
      formValue.value.groupIdsArr = [];
    }
    formValue.value.groupIds = formValue.value.groupIdsArr.join(',');
    e.preventDefault();
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        Edit(formValue.value)
          .then((_res) => {
            message.success('操作成功');
            setTimeout(() => {
              closeForm();
              formBtnLoading.value = false;
              emit('reloadTable');
            });
          })
          .catch((err) => {
            formBtnLoading.value = false;
          });
      } else {
        message.error('验证错误');
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
