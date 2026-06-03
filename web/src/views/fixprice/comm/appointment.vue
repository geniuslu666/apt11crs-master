<template>
  <!-- 预约订房 -->
  <div class="flex-column">
    <div class="flex-item pl-2 pr-2" style="overflow-x: hidden">
      <div class="p-3 f16 fw c333">基础信息</div>
      <n-form ref="formRef" inline :label-width="80" :model="state.formjb" :rules="state.rules">
        <n-row gutter="12" class="p-3">
          <n-col :span="12">
            <n-form-item label="物业" path="user.name">
              <n-select v-model:value="state.formjb.value" filterable :options="state.optionswy" />
            </n-form-item>
          </n-col>
          <n-col :span="12">
            <n-form-item label="房型" path="user.name">
              <n-select v-model:value="state.formjb.value" filterable :options="state.optionswy" />
            </n-form-item>
          </n-col>
          <n-col :span="12">
            <n-form-item label="入住日期" path="name">
              <n-date-picker v-model:value="state.formjb.date" type="daterange" clearable />
            </n-form-item>
          </n-col>
          <n-col :span="6">
            <n-form-item label="入住时间" path="phone">
              <n-time-picker default-formatted-value="00:12:13" style="width: 100%" />
            </n-form-item>
          </n-col>
          <n-col :span="6">
            <n-form-item label="退房时间" path="phone">
              <n-time-picker :default-value="1183135260000" style="width: 100%" />
            </n-form-item>
          </n-col>
          <n-col :span="8">
            <n-form-item label="成人数量" path="user.name">
              <n-input-number v-model:value="state.formjb.inputNumberValue" />
            </n-form-item>
          </n-col>

          <n-col :span="8">
            <n-form-item label="婴儿数量" path="user.name">
              <n-input-number v-model:value="state.formjb.inputNumberValue" />
            </n-form-item>
          </n-col>

          <n-col :span="8">
            <n-form-item label="儿童数量" path="user.name">
              <n-input-number v-model:value="state.formjb.inputNumberValue" />
            </n-form-item>
          </n-col>
        </n-row>
      </n-form>
      <div class="xvxian mb-3 ml-3 mr-3"></div>

      <!-- 预定者信息 -->
      <div class="p-3 f16 fw c333">预定者信息</div>
      <n-form ref="formRef" inline :label-width="80" :model="state.formjb" :rules="state.rules">
        <n-row gutter="12" class="p-3">
          <n-col :span="12">
            <n-form-item label="姓名" path="user.name">
              <n-input v-model:value="state.formjb.value" type="text" placeholder="预定者姓名" />
            </n-form-item>
          </n-col>
          <n-col :span="12">
            <n-form-item label="电话" path="user.name">
              <n-input v-model:value="state.formjb.value" type="text" placeholder="预定者姓名" />
            </n-form-item>
          </n-col>

          <n-col :span="12">
            <n-form-item label="邮箱" path="user.name">
              <n-input v-model:value="state.formjb.value" type="text" placeholder="预定者姓名" />
            </n-form-item>
          </n-col>

          <n-col :span="12">
            <n-form-item label="国家" path="user.name">
              <n-input v-model:value="state.formjb.value" type="text" placeholder="预定者姓名" />
            </n-form-item>
          </n-col>
        </n-row>
      </n-form>
      <div class="xvxian mb-3 ml-3 mr-3"></div>

      <!-- 入住人信息 -->
      <div>
        <div class="p-3 f16 fw c333">入住人信息</div>
        <div class="p-3">
          <n-form ref="formRef" inline :label-width="80" :model="state.formjb" :rules="state.rules">
            <n-row gutter="12">
              <n-col :span="12">
                <n-form-item label="姓名" path="user.name">
                  <n-input
                    v-model:value="state.formjb.value"
                    type="text"
                    placeholder="预定者姓名"
                  />
                </n-form-item>
              </n-col>
              <n-col :span="12">
                <n-form-item label="电话" path="user.name">
                  <n-input
                    v-model:value="state.formjb.value"
                    type="text"
                    placeholder="预定者姓名"
                  />
                </n-form-item>
              </n-col>

              <n-col :span="12">
                <n-form-item label="邮箱" path="user.name">
                  <n-input
                    v-model:value="state.formjb.value"
                    type="text"
                    placeholder="预定者姓名"
                  />
                </n-form-item>
              </n-col>

              <n-col :span="12">
                <n-form-item label="国家" path="user.name">
                  <n-input
                    v-model:value="state.formjb.value"
                    type="text"
                    placeholder="预定者姓名"
                  />
                </n-form-item>
              </n-col>
            </n-row>
          </n-form>
          <div>
            <n-button type="primary" size="small" ghost>
              <template #icon>
                <n-icon>
                  <FileAddOutlined />
                </n-icon>
              </template>
              新增入住人
            </n-button>
          </div>
        </div>
        <div class="xvxian mb-3 ml-3 mr-3"></div>
      </div>
    </div>
    <div class="drawerfooter">
      <div class="flex-row">
        <div class="flex-item text-l"></div>
        <div>
          <n-button class="mr-2">取消</n-button> <n-button type="primary"> 提交订单 </n-button></div
        >
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
  import { reactive, onMounted } from 'vue';
  import { FileAddOutlined } from '@vicons/antd';
  const state = reactive({
    selectdays: {},
    formjb: { value: '', date: [1183135260000, 1716962533623] },
    optionswy: [
      {
        label: "Everybody's Got Something to Hide Except Me and My Monkey",
        value: 'song0',
        disabled: true,
      },
    ],
    rules: {
      user: {
        name: {
          required: true,
          message: '请选择物业信息',
          trigger: 'blur',
        },
        age: {
          required: true,
          message: '请输入年龄',
          trigger: ['input', 'blur'],
        },
      },
      phone: {
        required: true,
        message: '请输入电话号码',
        trigger: ['input'],
      },
    },
  });
  // 对象写法并校验
  const props = defineProps({
    selectdays: Object,
  });
  onMounted(() => {
    state.selectdays = props.selectdays;
  });
</script>

<style lang="less" scoped></style>
