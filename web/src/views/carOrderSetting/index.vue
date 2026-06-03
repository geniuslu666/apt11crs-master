<template>
  <div>
    <n-spin :show="show" description="请稍候...">
      <n-card :bordered="false" :header-style="{
                    padding: '20px',
                  }">
        <template #header>
          <text style="font-weight: 500;font-size: 18px;color: #3D3D3D;line-height: 25px;">预约设置</text>
        </template>
        <n-form
          ref="formRef"
          :model="formValue"
          :rules="rules"
          :label-placement="settingStore.isMobile ? 'top' : 'left'"
          :label-width="150"
          class="py-4"
          style="padding-top: 0"
        >
          <n-grid cols="1 s:1 m:1 l:1 xl:1 2xl:1" responsive="screen">
            <n-gi span="1">
              <n-form-item label="预约日期" path="orderTimeType">
                <n-radio-group v-model:value="formValue.orderTimeType" name="orderTimeType">
                  <n-space>
                    <n-radio :value="1">
                      每天
                    </n-radio>
                    <n-radio :value="2">
                      自定义
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
              <n-form-item label=" " path="orderTimeWeekArr" v-if="formValue.orderTimeType == 2">
                <n-checkbox-group v-model:value="formValue.orderTimeWeekArr">
                  <n-space>
                    <n-checkbox
                      v-for="item in weekList"
                      :key="item.value"
                      :value="item.value"
                      :label="item.label"
                    />
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item :label="index == 0 ? '可预约时段' : ' '" :path="'orderTime'+index" v-for="(item, index) in orderTimeFormArr" :key="index">
                <n-time-picker placeholder="开始时间" format="HH:mm" v-model:formatted-value="item.startTime" />
                <span style="margin: 0 5px">~</span>
                <n-time-picker placeholder="结束时间" format="HH:mm" v-model:formatted-value="item.endTime" />
                <n-button text type="success" style="margin-left: 12px" attr-type="button" @click="addItem" v-if="index == 0">
                  <template #icon>
                    <n-icon>
                      <PlusCircleOutlined />
                    </n-icon>
                  </template>
                </n-button>
                <n-button text type="error" style="margin-left: 12px" attr-type="button" @click="removeItem(index)" v-if="index > 0">
                  <template #icon>
                    <n-icon>
                      <MinusCircleOutlined />
                    </n-icon>
                  </template>
                </n-button>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预约时段间隔" path="timeDuration">
                <n-radio-group v-model:value="formValue.timeDuration" name="timeDuration">
                  <n-radio-button
                    v-for="item in timeDurationArr"
                    :key="item.value"
                    :value="item.value"
                    :label="item.label"
                  />
                </n-radio-group>
                <n-input-group-label>分钟</n-input-group-label>
              </n-form-item>
            </n-gi>

            <n-gi span="1">
              <n-form-item label="提前预约" path="advanceOrderHour" style="margin-bottom: 24px">
                <n-input-group>
                  <n-input-group-label>需至少提前</n-input-group-label>
                  <n-input-number placeholder="请输入小时" :min="0" :precision="0" v-model:value="formValue.advanceOrderHour" style="width: 100px" />
                  <n-input-group-label>小时</n-input-group-label>
                </n-input-group>
                <template #feedback>
                  填0则当前可约
                </template>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="最长预约" path="maxOrderDay">
                <n-input-group>
                  <n-input-group-label>最长可预约</n-input-group-label>
                  <n-input-number placeholder="请输入天数" :min="0" :precision="0" :show-button="false" v-model:value="formValue.maxOrderDay" style="width: 100px" />
                  <n-input-group-label>天后</n-input-group-label>
                </n-input-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="派单模式" path="dispatchMode">
                <n-radio-group v-model:value="formValue.dispatchMode" name="dispatchMode">
                  <n-space>
                    <n-radio :value="1">
                      人工接单
                    </n-radio>
                    <n-radio :value="2">
                      系统自动指定
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="预约确定模式" path="orderConfirmType">
                <n-radio-group v-model:value="formValue.orderConfirmType" name="orderConfirmType">
                  <n-space>
                    <n-radio :value="1">
                      手动确认
                    </n-radio>
                    <n-radio :value="2">
                      自动确认
                    </n-radio>
                    <n-radio :value="3">
                      同时支持
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
              <n-form-item label="预计确认时长" path="orderConfirmTime" v-if="formValue.orderConfirmType == 1">
                <n-input-group>
                  <n-input-number placeholder="请输入" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmTime" style="width: 100px" />
                  <n-input-group-label>分钟</n-input-group-label>
                </n-input-group>
              </n-form-item>
              <template v-if="formValue.orderConfirmType == 3 && formValue.orderConfirmDayType == 1">
                <n-form-item label=" " path="orderConfirm">
                  <n-input-group>
                    <n-input-group-label>预定</n-input-group-label>
                    <n-input-number placeholder="请输入" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmBeforeDays" style="width: 100px" />
                    <n-input-group-label>天内手动确认，预计确认时长</n-input-group-label>
                    <n-input-number placeholder="确认时长" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmTime" style="width: 100px" />
                    <n-input-group-label>分钟</n-input-group-label>
                  </n-input-group>
                </n-form-item>
                <n-form-item label=" " path="orderConfirm">
                  <n-input-group>
                    <n-input-group-label>大于{{formValue.orderConfirmBeforeDays}}天自动确认</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </template>
              <template v-if="formValue.orderConfirmType == 3 && formValue.orderConfirmDayType == 2">
                <n-form-item label=" " path="orderConfirm">
                  <n-input-group>
                    <n-input-group-label>预定</n-input-group-label>
                    <n-input-number placeholder="请输入" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmBeforeDays" style="width: 100px" />
                    <n-input-group-label>天前自动确认</n-input-group-label>
                  </n-input-group>
                </n-form-item>
                <n-form-item label=" " path="orderConfirm">
                  <n-input-group>
                    <n-input-group-label>{{formValue.orderConfirmBeforeDays}}天之后手动确认，预计确认时长</n-input-group-label>
                    <n-input-number placeholder="确认时长" :min="0" :precision="0" :show-button="false" v-model:value="formValue.orderConfirmTime" style="width: 100px" />
                    <n-input-group-label>分钟</n-input-group-label>
                  </n-input-group>
                </n-form-item>
              </template>
            </n-gi>
            <n-gi span="1">
              <n-form-item label="每日限定" path="maxOneDayOrderNumJson">
                <!-- 每日接机送机限制设置 -->
                <div style="display: flex; align-items: center; gap: 16px; padding: 12px 16px; background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%); border-radius: 10px; border: 1px solid #e2e8f0; max-width: 400px;">
                  <!-- 接机每日限制 -->
                  <div style="display: flex; align-items: center; gap: 8px;">
                    <div style="display: flex; align-items: center; justify-content: center; width: 24px; height: 24px; background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%); border-radius: 6px; box-shadow: 0 2px 6px rgba(59, 130, 246, 0.25);">
                      <span style="color: white; font-size: 11px; font-weight: 700;">接</span>
                    </div>
                    <n-input-number
                      placeholder=""
                      :min="0"
                      :precision="0"
                      :show-button="false"
                      v-model:value="dailyLimits.pickupNum"
                      size="small"
                      style="width: 70px;"
                    />
                    <span style="color: #64748b; font-size: 12px; font-weight: 500;">单</span>
                  </div>

                  <!-- 分隔线 -->
                  <div style="width: 1px; height: 20px; background: linear-gradient(to bottom, transparent, #cbd5e1, transparent);"></div>

                  <!-- 送机每日限制 -->
                  <div style="display: flex; align-items: center; gap: 8px;">
                    <div style="display: flex; align-items: center; justify-content: center; width: 24px; height: 24px; background: linear-gradient(135deg, #10b981 0%, #059669 100%); border-radius: 6px; box-shadow: 0 2px 6px rgba(16, 185, 129, 0.25);">
                      <span style="color: white; font-size: 11px; font-weight: 700;">送</span>
                    </div>
                    <n-input-number
                      placeholder=""
                      :min="0"
                      :precision="0"
                      :show-button="false"
                      v-model:value="dailyLimits.dropoffNum"
                      size="small"
                      style="width: 70px;"
                    />
                    <span style="color: #64748b; font-size: 12px; font-weight: 500;">单</span>
                  </div>
                </div>
                <template #feedback>
                  填0不限制
                </template>
              </n-form-item>
            </n-gi>
            <n-gi span="1" style="margin-top: 24px">
              <n-form-item label="时间段限制" path="maxTimeOrderOpen">
                <n-radio-group v-model:value="formValue.maxTimeOrderOpen" name="maxTimeOrderOpen">
                  <n-space>
                    <n-radio :value="1">
                      开启
                    </n-radio>
                    <n-radio :value="2">
                      关闭
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
              <n-form-item :label="index == 0 ? '时间段' : ' '" :path="'orderTime1'+index" v-for="(item, index) in orderTimeForm1Arr" :key="index" v-if="formValue.maxTimeOrderOpen == 1">
                <n-time-picker placeholder="开始时间" format="HH:mm" v-model:formatted-value="item.startTime" />
                <span style="margin: 0 5px">~</span>
                <n-time-picker placeholder="结束时间" format="HH:mm" v-model:formatted-value="item.endTime" />
                <n-input-group style="margin-left: 5px">
                  <n-input-number placeholder="请输入单数" :min="0" :precision="0" :show-button="false" v-model:value="item.num" style="width: 100px" />
                  <n-input-group-label style="margin-right: 12px">单</n-input-group-label>
                  <n-button text type="success" style="margin-left: 12px" attr-type="button" @click="addItem1" v-if="index == 0">
                    <template #icon>
                      <n-icon>
                        <PlusCircleOutlined />
                      </n-icon>
                    </template>
                  </n-button>
                  <n-button text type="error" style="margin-left: 12px" attr-type="button" @click="removeItem1(index)" v-if="index > 0">
                    <template #icon>
                      <n-icon>
                        <MinusCircleOutlined />
                      </n-icon>
                    </template>
                  </n-button>
                </n-input-group>

              </n-form-item>
            </n-gi>
            <n-gi span="1" style="margin-top: 24px">
              <n-form-item label="日期限制" path="maxDateTimeOrderOpen" :show-feedback="false">
                <n-radio-group v-model:value="formValue.maxDateTimeOrderOpen" name="maxDateTimeOrderOpen">
                  <n-space>
                    <n-radio :value="1">
                      开启
                    </n-radio>
                    <n-radio :value="2">
                      关闭
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>

              <!-- 新的多日期时间段设置 -->
              <div v-if="formValue.maxDateTimeOrderOpen == 1" style="margin-top: 20px;margin-left: 150px;max-width: 800px;">
                <n-card
                  v-for="(dateItem, dateIndex) in dateTimeSettingArr"
                  :key="dateIndex"
                  :bordered="true"
                  size="small"
                  style="margin-bottom: 16px; border: 1px solid #e0e0e6; border-radius: 8px;"
                  :header-style="{ padding: '12px 16px', backgroundColor: '#fafafa', borderBottom: '1px solid #e0e0e6' }"
                  :content-style="{ padding: '16px' }"
                >
                  <template #header>
                    <div style="display: flex; align-items: center; justify-content: space-between;">
                      <div style="display: flex; align-items: center; gap: 12px;">
                        <n-icon size="16" color="#18a058">
                          <svg viewBox="0 0 24 24" fill="currentColor">
                            <path d="M19 3h-1V1h-2v2H8V1H6v2H5c-1.11 0-1.99.9-1.99 2L3 19c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V8h14v11zM7 10h5v5H7z"/>
                          </svg>
                        </n-icon>
                        <span style="font-weight: 500; color: #333;">日期 {{ dateIndex + 1 }}</span>
                      </div>
                      <n-button
                        text
                        type="error"
                        size="small"
                        @click="removeDateSetting(dateIndex)"
                        v-if="dateTimeSettingArr.length > 1"
                        style="padding: 4px;"
                      >
                        <template #icon>
                          <n-icon size="14">
                            <MinusCircleOutlined />
                          </n-icon>
                        </template>
                        删除日期
                      </n-button>
                    </div>
                  </template>

                  <div>
                    <!-- 日期选择 -->
                    <n-form-item :label="`日期选择`" :show-feedback="false" style="margin-bottom: 16px;">
                      <n-date-picker
                        placeholder="选择日期"
                        format="yyyy-MM-dd"
                        v-model:formatted-value="dateItem.date"
                        style="width: 200px;"
                      />
                    </n-form-item>

                    <!-- 时间段设置 -->
                    <div style="border-top: 1px solid #f0f0f0; padding-top: 16px;">
                      <div style="margin-bottom: 12px; display: flex; align-items: center; justify-content: space-between;">
                        <span style="font-weight: 500; color: #666; font-size: 14px;">时间段设置</span>
                        <n-button
                          text
                          type="success"
                          size="small"
                          @click="addTimeSlot(dateIndex)"
                          style="padding: 4px 8px;"
                        >
                          <template #icon>
                            <n-icon size="14">
                              <PlusCircleOutlined />
                            </n-icon>
                          </template>
                          添加时间段
                        </n-button>
                      </div>

                      <div
                        v-for="(timeSlot, timeIndex) in dateItem.timeSlots"
                        :key="timeIndex"
                        style="display: flex; align-items: center; gap: 8px; margin-bottom: 12px; padding: 12px; background-color: #f8f9fa; border-radius: 6px; border: 1px solid #e9ecef;"
                      >
                        <n-time-picker
                          placeholder="开始时间"
                          format="HH:mm"
                          v-model:formatted-value="timeSlot.startTime"
                          size="small"
                          style="width: 140px;"
                        />
                        <span style="color: #666; font-weight: 500;">~</span>
                        <n-time-picker
                          placeholder="结束时间"
                          format="HH:mm"
                          v-model:formatted-value="timeSlot.endTime"
                          size="small"
                          style="width: 140px;"
                        />
                        <!-- 车型限制设置 -->
                        <div style="display: flex; flex-direction: column; gap: 12px; padding: 12px; background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%); border-radius: 8px; border: 1px solid #e2e8f0; min-width: 380px; flex: 1;">
                          <!-- 车型限制开关 -->
                          <div style="display: flex; align-items: center; gap: 8px; padding-bottom: 8px; border-bottom: 1px solid #e2e8f0;">
                            <n-checkbox v-model:checked="timeSlot.hasCarTypes" size="small">
                              <span style="color: #475569; font-size: 13px; font-weight: 500;">按车型限制</span>
                            </n-checkbox>
                          </div>

                          <!-- 不限制车型时显示总数量 -->
                          <div v-if="!timeSlot.hasCarTypes" style="display: flex; align-items: center; gap: 12px;">
                            <!-- 接机设置 -->
                            <div style="display: flex; align-items: center; gap: 6px;">
                              <div style="display: flex; align-items: center; justify-content: center; width: 20px; height: 20px; background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%); border-radius: 4px; box-shadow: 0 2px 4px rgba(59, 130, 246, 0.2);">
                                <span style="color: white; font-size: 10px; font-weight: 600;">接</span>
                              </div>
                              <n-input-number
                                placeholder=""
                                :min="0"
                                :precision="0"
                                :show-button="false"
                                v-model:value="timeSlot.pickupNum"
                                size="small"
                                style="width: 70px;"
                              />
                              <span style="color: #64748b; font-size: 12px; font-weight: 500;">单</span>
                            </div>

                            <!-- 分隔线 -->
                            <div style="width: 1px; height: 16px; background: linear-gradient(to bottom, transparent, #cbd5e1, transparent);"></div>

                            <!-- 送机设置 -->
                            <div style="display: flex; align-items: center; gap: 6px;">
                              <div style="display: flex; align-items: center; justify-content: center; width: 20px; height: 20px; background: linear-gradient(135deg, #10b981 0%, #059669 100%); border-radius: 4px; box-shadow: 0 2px 4px rgba(16, 185, 129, 0.2);">
                                <span style="color: white; font-size: 10px; font-weight: 600;">送</span>
                              </div>
                              <n-input-number
                                placeholder=""
                                :min="0"
                                :precision="0"
                                :show-button="false"
                                v-model:value="timeSlot.dropoffNum"
                                size="small"
                                style="width: 70px;"
                              />
                              <span style="color: #64748b; font-size: 12px; font-weight: 500;">单</span>
                            </div>
                          </div>

                          <!-- 按车型限制时显示车型列表 -->
                          <div v-if="timeSlot.hasCarTypes" style="display: flex; flex-direction: column; gap: 8px;">
                            <!-- 添加车型按钮 -->
                            <div style="display: flex; justify-content: space-between; align-items: center;">
                              <span style="color: #475569; font-size: 12px; font-weight: 500;">车型设置</span>
                              <n-button
                                text
                                type="primary"
                                size="small"
                                @click="addVehicleType(dateIndex, timeIndex)"
                                style="padding: 2px 6px;"
                              >
                                <template #icon>
                                  <n-icon size="12">
                                    <PlusCircleOutlined />
                                  </n-icon>
                                </template>
                                添加车型
                              </n-button>
                            </div>

                            <!-- 车型列表 -->
                            <div
                              v-for="(vehicle, vehicleIndex) in timeSlot.carTypes"
                              :key="vehicleIndex"
                              style="display: flex; align-items: center; gap: 8px; padding: 8px; background: white; border-radius: 6px; border: 1px solid #e2e8f0;"
                            >
                              <!-- 车型选择 -->
                              <n-select
                                placeholder="选择车型"
                                v-model:value="vehicle.id"
                                :options="vehicleTypeOptions"
                                label-field="name"
                                value-field="id"
                                size="small"
                                style="width: 120px;"
                                @update:value="(value) => onCarTypeChange(dateIndex, timeIndex, vehicleIndex, value)"
                              />

                              <!-- 接机数量 -->
                              <div style="display: flex; align-items: center; gap: 4px;">
                                <div style="display: flex; align-items: center; justify-content: center; width: 16px; height: 16px; background: linear-gradient(135deg, #3b82f6 0%, #1d4ed8 100%); border-radius: 3px;">
                                  <span style="color: white; font-size: 8px; font-weight: 600;">接</span>
                                </div>
                                <n-input-number
                                  placeholder=""
                                  :min="0"
                                  :precision="0"
                                  :show-button="false"
                                  v-model:value="vehicle.pickupNum"
                                  size="small"
                                  style="width: 60px;"
                                />
                              </div>

                              <!-- 送机数量 -->
                              <div style="display: flex; align-items: center; gap: 4px;">
                                <div style="display: flex; align-items: center; justify-content: center; width: 16px; height: 16px; background: linear-gradient(135deg, #10b981 0%, #059669 100%); border-radius: 3px;">
                                  <span style="color: white; font-size: 8px; font-weight: 600;">送</span>
                                </div>
                                <n-input-number
                                  placeholder=""
                                  :min="0"
                                  :precision="0"
                                  :show-button="false"
                                  v-model:value="vehicle.dropoffNum"
                                  size="small"
                                  style="width: 60px;"
                                />
                              </div>

                              <!-- 删除车型按钮 -->
                              <n-button
                                text
                                type="error"
                                size="small"
                                @click="removeVehicleType(dateIndex, timeIndex, vehicleIndex)"
                                style="padding: 2px;"
                              >
                                <template #icon>
                                  <n-icon size="12">
                                    <MinusCircleOutlined />
                                  </n-icon>
                                </template>
                              </n-button>
                            </div>

                            <!-- 无车型提示 -->
                            <div v-if="timeSlot.carTypes.length === 0" style="text-align: center; padding: 12px; color: #9ca3af; font-size: 12px;">
                              请添加车型设置
                            </div>
                          </div>
                        </div>
                        <n-button
                          text
                          type="error"
                          size="small"
                          @click="removeTimeSlot(dateIndex, timeIndex)"
                          v-if="dateItem.timeSlots.length > 1"
                          style="margin-left: 8px; padding: 4px;"
                        >
                          <template #icon>
                            <n-icon size="14">
                              <MinusCircleOutlined />
                            </n-icon>
                          </template>
                        </n-button>
                      </div>
                    </div>
                  </div>
                </n-card>

                <!-- 添加新日期按钮 -->
                <div style="text-align: center; margin-top: 16px;">
                  <n-button
                    dashed
                    type="primary"
                    @click="addDateSetting"
                    style="width: 100%; height: 48px; border: 2px dashed #18a058;"
                  >
                    <template #icon>
                      <n-icon size="18">
                        <PlusCircleOutlined />
                      </n-icon>
                    </template>
                    添加新日期设置
                  </n-button>
                </div>
              </div>
            </n-gi>
            <n-gi span="1" style="margin-top: 24px">
              <n-form-item label="深夜费" path="nightFeeEnabled">
                <n-radio-group v-model:value="formValue.nightFeeEnabled" name="nightFeeEnabled">
                  <n-space>
                    <n-radio :value="1">
                      开启
                    </n-radio>
                    <n-radio :value="2">
                      关闭
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
              
              <!-- 深夜费配置区域 -->
              <div v-if="formValue.nightFeeEnabled == 1" style="margin-left: 150px; max-width: 800px; padding: 20px; background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%); border-radius: 12px; border: 1px solid #e2e8f0; box-shadow: 0 2px 8px rgba(0, 0, 0, 0.04);">
                <!-- 服务类型选择 -->
                <n-form-item label=" " label-width="0" :label-style="{padding:'0px'}" :show-feedback="false" path="nightFeeServiceTypesStr" style="margin-bottom: 16px;">
                  <div style="display: flex; align-items: center; gap: 12px;">
                    <span style="color: #374151; font-size: 14px; font-weight: 500; min-width: 60px;">服务类型</span>
                    <n-checkbox-group v-model:value="nightFeeServiceTypesArray">
                      <n-space>
                        <n-checkbox value="pickup" label="接机" />
                        <n-checkbox value="dropoff" label="送机" />
                      </n-space>
                    </n-checkbox-group>
                  </div>
                </n-form-item>

                <!-- 统一价格选项 -->
                <n-form-item label=" " label-width="0" :label-style="{padding:'0px'}" :show-feedback="false" path="nightFeeUnifiedPrice" style="margin-bottom: 16px;">
                  <div style="display: flex; align-items: center; gap: 12px;">
                    <n-switch
                      v-model:value="formValue.nightFeeUnifiedPrice"
                      :checked-value="1"
                      :unchecked-value="2"
                      size="small"
                    />
                    <span style="color: #374151; font-size: 14px; font-weight: 500;">统一价格</span>
                  </div>
                </n-form-item>
                
                <!-- 统一价格输入框 -->
                <n-form-item 
                  label=" " 
                  label-width="0"
                  :label-style="{padding:'0px'}"
                  :show-feedback="false"
                  path="nightFeeUnifiedAmount" 
                  v-if="formValue.nightFeeUnifiedPrice == 1"
                  style="margin-bottom: 20px;"
                >
                  <n-input-group style="max-width: 200px;">
                    <n-input-group-label>价格</n-input-group-label>
                    <n-input-number 
                      placeholder="0" 
                      :min="0" 
                      :precision="0" 
                      :show-button="false" 
                      v-model:value="formValue.nightFeeUnifiedAmount" 
                      style="width: 120px" 
                    />
                    <n-input-group-label>JPY</n-input-group-label>
                  </n-input-group>
                </n-form-item>

                <!-- 深夜费时间段配置 -->
                <n-form-item :label="index == 0 ? '时间段' : ' '" label-width="60" label-align="left" :path="'nightFeeTime'+index" v-for="(timeSlot, index) in nightFeeTimeSlots" :key="index">
                  <div style="display: flex; align-items: center; flex-wrap: wrap; gap: 8px; width: 100%;">
                    <!-- 时间选择器组 -->
                    <div style="display: flex; align-items: center; gap: 8px; min-width: 280px;">
                      <n-time-picker placeholder="开始时间" format="HH:mm" v-model:formatted-value="timeSlot.startTime" style="width: 130px;"/>
                      <span style="margin: 0 4px; color: #666;">~</span>
                      <n-time-picker placeholder="结束时间" format="HH:mm" v-model:formatted-value="timeSlot.endTime" style="width: 130px;"/>
                    </div>
                    
                    <!-- 价格输入框（仅在未勾选统一价格时显示） -->
                    <div v-if="formValue.nightFeeUnifiedPrice == 2" style="display: flex; align-items: center; gap: 4px; min-width: 140px;">
                      <n-input-group style="max-width: 140px;">
                        <n-input-number placeholder="请输入价格" :min="0" :precision="0" :show-button="false" v-model:value="timeSlot.price" style="width: 100px" />
                        <n-input-group-label>JPY</n-input-group-label>
                      </n-input-group>
                    </div>
                    
                    <!-- 操作按钮 -->
                    <div style="display: flex; align-items: center; gap: 8px; margin-left: 12px;">
                      <n-button text type="success" attr-type="button" @click="addNightFeeTimeSlot" v-if="index == 0">
                        <template #icon>
                          <n-icon>
                            <PlusCircleOutlined />
                          </n-icon>
                        </template>
                      </n-button>
                      <n-button text type="error" attr-type="button" @click="removeNightFeeTimeSlot(index)" v-if="index > 0">
                        <template #icon>
                          <n-icon>
                            <MinusCircleOutlined />
                          </n-icon>
                        </template>
                      </n-button>
                    </div>
                  </div>
                </n-form-item>
                
                <!-- 时间段说明提示 -->
                <div style="margin-top: 8px; padding-left: 60px;">
                  <span style="color: #e74c3c; font-size: 12px;">
                    注意：时间段包含开始时间，不包含结束时间。例如：00:00~06:00 表示 00:00≤时间&lt;06:00
                  </span>
                </div>
              </div>
            </n-gi>
            <n-gi span="1" style="margin-top: 24px;">
              <n-form-item label="派单限制" path="orderDispatch">
                <n-radio-group v-model:value="formValue.orderDispatch" name="orderDispatch">
                  <n-space>
                    <n-radio :value="1">
                      不限制
                    </n-radio>
                    <n-radio :value="2">
                      限制：出车中、休息中
                    </n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
            </n-gi>
          </n-grid>
          <div style="text-align: center">
            <n-space justify="center">
              <n-button type="info" :loading="formBtnLoading" @click="formSubmit">
                确定
              </n-button>
            </n-space>
          </div>
        </n-form>
      </n-card>
    </n-spin>
  </div>
</template>

<script lang="ts" setup>
import {ref, onMounted, reactive} from 'vue';
  import { useMessage } from 'naive-ui';
  import { getConfig, updateConfig } from '@/api/sys/config';
  import {useProjectSettingStore} from "@/store/modules/projectSetting";
import {MinusCircleOutlined, PlusCircleOutlined} from "@vicons/antd";
import { List as CarTypeList } from '@/api/carCarType';

  const rules = ref({});
  const group = ref('carordersetting');
  const show = ref(false);
  const formRef: any = ref(null);
  const message = useMessage();
  const formBtnLoading = ref(false);
  const settingStore = useProjectSettingStore();

  const formValue = ref({
    orderTimeType: 1,
    orderTimeWeekArr: [],
    orderTimeWeek: '',
    timeDuration: 30,
    orderTimeForm: '',
    orderTimeForm1: '',
    orderTimeForm2: '',
    advanceOrderHour: null,
    maxOrderDay: null,
    dispatchMode: 1,
    orderConfirmDayType: 1,
    orderConfirmType: 1,
    orderConfirmBeforeDays: null,
    orderConfirmTime: null,
    maxOneDayOrderNumJson: '', // JSON格式存储：{"pickupNum": 0, "dropoffNum": 0}
    maxTimeOrderOpen: 2,
    maxDateTimeOrderOpen: 2,
    maxDateTimeOrderDate: null,
    maxDateTimeOrderDateTimeJson: '', // 新增：存储多日期时间段设置的JSON字符串
    
    // 深夜费相关配置
    nightFeeEnabled: 2, // 1-开启，2-关闭
    nightFeeServiceTypesStr: '', // 服务类型字符串：'pickup,dropoff'
    nightFeeUnifiedPrice: 2, // 统一价格：1-开启，2-关闭
    nightFeeUnifiedAmount: null, // 统一价格金额
    nightFeeTimeSlots: '', // JSON格式存储深夜费时间段配置

    orderDispatch: 1,
  });
  const weekList = ref([
    {
      label: '周一',
      value: 1,
    },
    {
      label: '周二',
      value: 2,
    },
    {
      label: '周三',
      value: 3,
    },
    {
      label: '周四',
      value: 4,
    },
    {
      label: '周五',
      value: 5,
    },
    {
      label: '周六',
      value: 6,
    },
    {
      label: '周日',
      value: 7,
    },
  ])
  const orderTimeFormArr = ref([{ startTime: null, endTime: null }]);
  const orderTimeForm1Arr = ref([{ startTime: null, endTime: null, num: null }]);
  const orderTimeForm2Arr = ref([{ startTime: null, endTime: null, num: null }]);

  // 车型列表数据
  interface CarTypeOption {
    id: number;
    name: string;
  }
  const vehicleTypeOptions = ref<CarTypeOption[]>([]);

  async function loadCarTypeList(){
    let carTypeArrListOrg = await CarTypeList({
      status: 1,
      Pagination: false
    })
    vehicleTypeOptions.value = carTypeArrListOrg.list
  }

  // 每日限制的响应式数据
  const dailyLimits = ref({
    pickupNum: null,
    dropoffNum: null,
  });

  // 新的多日期时间段设置数据
  const dateTimeSettingArr = ref<Array<{
    date: string|null;
    timeSlots: Array<{
      startTime: string|null;
      endTime: string|null;
      pickupNum: number;
      dropoffNum: number;
      hasCarTypes: boolean;
      carTypes: Array<{
        id: number|null;
        name: string|null;
        pickupNum: number;
        dropoffNum: number;
      }>;
    }>;
  }>>([]);

  // 深夜费时间段配置数据
  const nightFeeTimeSlots = ref<Array<{
    startTime: string|null;
    endTime: string|null;
    price: number|null;
  }>>([{ startTime: null, endTime: null, price: null }]);

  // 深夜费服务类型数组（用于前端UI绑定）
  const nightFeeServiceTypesArray = ref<string[]>([]);

  const timeDurationArr = reactive([
    {
      label: '15',
      value: 15
    },
    {
      label: '30',
      value: 30
    },
    {
      label: '60',
      value: 60
    },
    {
      label: '90',
      value: 90
    },
    {
      label: '120',
      value: 120
    },
  ])

  const removeItem = (index: number) => {
    orderTimeFormArr.value.splice(index, 1)
  }

  const addItem = () => {
    orderTimeFormArr.value.push({ startTime: null, endTime: null })
  }

const removeItem1 = (index: number) => {
  orderTimeForm1Arr.value.splice(index, 1)
}

const addItem1 = () => {
  orderTimeForm1Arr.value.push({ startTime: null, endTime: null, num: null })
}

// 这些方法保留但不使用，因为原有的注释部分可能还会用到
// const removeItem2 = (index: number) => {
//   orderTimeForm2Arr.value.splice(index, 1)
// }

// const addItem2 = () => {
//   orderTimeForm2Arr.value.push({ startTime: null, endTime: null, num: null })
// }

// 深夜费时间段配置方法
const addNightFeeTimeSlot = () => {
  nightFeeTimeSlots.value.push({
    startTime: null,
    endTime: null,
    price: null
  })
}

const removeNightFeeTimeSlot = (index: number) => {
  nightFeeTimeSlots.value.splice(index, 1)
}

// 新的多日期时间段设置方法
const addDateSetting = () => {
  dateTimeSettingArr.value.push({
    date: null,
    timeSlots: [
      {
        startTime: null,
        endTime: null,
        pickupNum: 0,
        dropoffNum: 0,
        hasCarTypes: false,
        carTypes: []
      }
    ]
  })
}

const removeDateSetting = (dateIndex: number) => {
  dateTimeSettingArr.value.splice(dateIndex, 1)
}

const addTimeSlot = (dateIndex: number) => {
  dateTimeSettingArr.value[dateIndex].timeSlots.push({
    startTime: null,
    endTime: null,
    pickupNum: 0,
    dropoffNum: 0,
    hasCarTypes: false,
    carTypes: []
  })
}

const removeTimeSlot = (dateIndex: number, timeIndex: number) => {
  dateTimeSettingArr.value[dateIndex].timeSlots.splice(timeIndex, 1)
}

// 车型管理函数
const addVehicleType = (dateIndex: number, timeIndex: number) => {
  dateTimeSettingArr.value[dateIndex].timeSlots[timeIndex].carTypes.push({
    id: null,
    name: null,
    pickupNum: 0,
    dropoffNum: 0
  })
}

const removeVehicleType = (dateIndex: number, timeIndex: number, vehicleIndex: number) => {
  dateTimeSettingArr.value[dateIndex].timeSlots[timeIndex].carTypes.splice(vehicleIndex, 1)
}

// 车型选择变化处理函数
const onCarTypeChange = (dateIndex: number, timeIndex: number, vehicleIndex: number, carTypeId: number | null) => {
  const carType = dateTimeSettingArr.value[dateIndex].timeSlots[timeIndex].carTypes[vehicleIndex]
  carType.id = carTypeId

  // 根据选择的车型ID找到对应的车型名称
  if (carTypeId) {
    const selectedCarType = vehicleTypeOptions.value.find(option => option.id === carTypeId)
    carType.name = selectedCarType ? selectedCarType.name : null
  } else {
    carType.name = null
  }
}

  function formSubmit() {
    formBtnLoading.value = true;
    formRef.value.validate((errors) => {
      if (!errors) {
        formValue.value.orderTimeWeek = formValue.value.orderTimeWeekArr ? formValue.value.orderTimeWeekArr.join(',') : ''
        formValue.value.orderTimeForm = JSON.stringify(orderTimeFormArr.value)
        formValue.value.orderTimeForm1 = JSON.stringify(orderTimeForm1Arr.value)
        formValue.value.orderTimeForm2 = JSON.stringify(orderTimeForm2Arr.value)

        // 计算车型数量总和并更新到timeSlot的pickupNum和dropoffNum
        dateTimeSettingArr.value.forEach(dateItem => {
          dateItem.timeSlots.forEach(timeSlot => {
            if (timeSlot.hasCarTypes && timeSlot.carTypes && timeSlot.carTypes.length > 0) {
              // 计算所有车型的接机数量总和
              timeSlot.pickupNum = timeSlot.carTypes.reduce((total, carType) => {
                return total + (carType.pickupNum || 0);
              }, 0);

              // 计算所有车型的送机数量总和
              timeSlot.dropoffNum = timeSlot.carTypes.reduce((total, carType) => {
                return total + (carType.dropoffNum || 0);
              }, 0);
            }
          });
        });

        formValue.value.maxDateTimeOrderDateTimeJson = JSON.stringify(dateTimeSettingArr.value)
        formValue.value.maxOneDayOrderNumJson = JSON.stringify(dailyLimits.value)

        let hasErr = false;
        orderTimeFormArr.value.forEach((item) => {
          if(!item.startTime || !item.endTime){
            message.error('预约时段请填写完整');
            hasErr = true;
            return false;
          }
          // if((new Date('2024-01-01 ' + item.startTime).getTime()) >= (new Date('2024-01-01 ' + item.endTime).getTime())){
          //   message.error('预约时段结束时间应该大于开始时间');
          //   hasErr = true;
          //   return false;
          // }
        })

        // 验证深夜费设置
        if (formValue.value.nightFeeEnabled == 1) {
          // 检查服务类型设置
          if (!nightFeeServiceTypesArray.value || nightFeeServiceTypesArray.value.length === 0) {
            message.error('开启深夜费后，请至少选择一种服务类型（接机或送机）');
            hasErr = true;
          }

          // 检查统一价格设置
          if (formValue.value.nightFeeUnifiedPrice == 1 && (!formValue.value.nightFeeUnifiedAmount || formValue.value.nightFeeUnifiedAmount <= 0)) {
            message.error('开启统一价格后，请填写有效的统一价格金额');
            hasErr = true;
          }

          // 检查时间段设置
          if (nightFeeTimeSlots.value.length === 0) {
            message.error('开启深夜费后，请至少添加一个时间段');
            hasErr = true;
          } else {
            for (let i = 0; i < nightFeeTimeSlots.value.length; i++) {
              const timeSlot = nightFeeTimeSlots.value[i];

              // 检查时间是否填写完整
              if (!timeSlot.startTime || !timeSlot.endTime) {
                message.error(`深夜费第${i + 1}个时间段：请填写完整的开始时间和结束时间`);
                hasErr = true;
                break;
              }

              // 检查时间段的逻辑性（开始时间应小于结束时间）
              const startTime = new Date('2024-01-01 ' + timeSlot.startTime).getTime();
              const endTime = new Date('2024-01-01 ' + timeSlot.endTime).getTime();

              if (startTime >= endTime) {
                message.error(`深夜费第${i + 1}个时间段：结束时间应该大于开始时间`);
                hasErr = true;
                break;
              }

              // 如果未勾选统一价格，检查每个时间段的价格
              if (formValue.value.nightFeeUnifiedPrice == 2) {
                if (!timeSlot.price || timeSlot.price <= 0) {
                  message.error(`深夜费第${i + 1}个时间段：请填写有效的价格`);
                  hasErr = true;
                  break;
                }
              }
            }

            // 检查时间段是否有重叠
            if (!hasErr) {
              for (let i = 0; i < nightFeeTimeSlots.value.length; i++) {
                for (let j = i + 1; j < nightFeeTimeSlots.value.length; j++) {
                  const slot1 = nightFeeTimeSlots.value[i];
                  const slot2 = nightFeeTimeSlots.value[j];
                  
                  const start1 = new Date('2024-01-01 ' + slot1.startTime).getTime();
                  const end1 = new Date('2024-01-01 ' + slot1.endTime).getTime();
                  const start2 = new Date('2024-01-01 ' + slot2.startTime).getTime();
                  const end2 = new Date('2024-01-01 ' + slot2.endTime).getTime();

                  // 检查时间段重叠
                  if ((start1 < end2 && end1 > start2)) {
                    message.error(`深夜费第${i + 1}个时间段与第${j + 1}个时间段存在重叠，请调整时间`);
                    hasErr = true;
                    break;
                  }
                }
                if (hasErr) break;
              }
            }
          }
        }

        // 验证日期限制设置
        if (formValue.value.maxDateTimeOrderOpen == 1) {
          // 检查是否有日期限制设置
          if (dateTimeSettingArr.value.length === 0) {
            message.error('开启日期限制后，请至少添加一个日期设置');
            hasErr = true;
          } else {
            // 用于检查重复日期
            const dateSet = new Set();

            for (let i = 0; i < dateTimeSettingArr.value.length; i++) {
              const dateItem = dateTimeSettingArr.value[i];

              // 检查日期是否填写
              if (!dateItem.date) {
                message.error(`第${i + 1}个日期设置：请选择日期`);
                hasErr = true;
                break;
              }

              // 检查日期是否重复
              if (dateSet.has(dateItem.date)) {
                message.error(`第${i + 1}个日期设置：日期"${dateItem.date}"已存在，请选择其他日期`);
                hasErr = true;
                break;
              }
              dateSet.add(dateItem.date);

              // 检查时间段是否填写
              if (!dateItem.timeSlots || dateItem.timeSlots.length === 0) {
                message.error(`第${i + 1}个日期设置：请至少添加一个时间段`);
                hasErr = true;
                break;
              }

              // 检查每个时间段的完整性
              for (let j = 0; j < dateItem.timeSlots.length; j++) {
                const timeSlot = dateItem.timeSlots[j];

                if (!timeSlot.startTime || !timeSlot.endTime) {
                  message.error(`第${i + 1}个日期设置，第${j + 1}个时间段：请填写完整的开始时间和结束时间`);
                  hasErr = true;
                  break;
                }

                // 检查时间段的逻辑性（开始时间应小于结束时间）
                const startTime = new Date('2024-01-01 ' + timeSlot.startTime).getTime();
                const endTime = new Date('2024-01-01 ' + timeSlot.endTime).getTime();

                if (startTime >= endTime) {
                  message.error(`第${i + 1}个日期设置，第${j + 1}个时间段：结束时间应该大于开始时间`);
                  hasErr = true;
                  break;
                }

                // 检查车型限制设置
                if (timeSlot.hasCarTypes) {
                  // 如果启用了车型限制，检查是否添加了车型
                  if (!timeSlot.carTypes || timeSlot.carTypes.length === 0) {
                    message.error(`第${i + 1}个日期设置，第${j + 1}个时间段：启用车型限制后，请至少添加一个车型`);
                    hasErr = true;
                    break;
                  }

                  // 检查每个车型的设置
                  for (let k = 0; k < timeSlot.carTypes.length; k++) {
                    const carType = timeSlot.carTypes[k];

                    // 检查是否选择了车型
                    if (!carType.id) {
                      message.error(`第${i + 1}个日期设置，第${j + 1}个时间段，第${k + 1}个车型：请选择车型`);
                      hasErr = true;
                      break;
                    }

                    // 检查接机和送机数量是否至少有一个大于0
                    if (( carType.pickupNum === null || carType.pickupNum < 0) || ( carType.dropoffNum === null || carType.dropoffNum < 0)) {
                      message.error(`第${i + 1}个日期设置，第${j + 1}个时间段，第${k + 1}个车型：接机数量和送机数量需要设置大于等于0`);
                      hasErr = true;
                      break;
                    }
                  }

                  // 检查车型是否重复
                  const carTypeIds = timeSlot.carTypes.map(ct => ct.id).filter(id => id);
                  const uniqueCarTypeIds = new Set(carTypeIds);
                  if (carTypeIds.length !== uniqueCarTypeIds.size) {
                    message.error(`第${i + 1}个日期设置，第${j + 1}个时间段：存在重复的车型，请检查`);
                    hasErr = true;
                    break;
                  }
                } else {
                  // 如果未启用车型限制，检查总数量设置
                  if ((timeSlot.pickupNum === null || timeSlot.pickupNum < 0) || (timeSlot.dropoffNum === null || timeSlot.dropoffNum < 0)) {
                    message.error(`第${i + 1}个日期设置，第${j + 1}个时间段：接机数量和送机数量需要设置大于等于0`);
                    hasErr = true;
                    break;
                  }
                }
              }

              if (hasErr) break;
            }
          }
        }

        if(hasErr){
          formBtnLoading.value = false;
          return false;
        }

        // 处理深夜费数据
        if (formValue.value.nightFeeEnabled == 1) {
          // 将服务类型数组转换为逗号分隔的字符串
          formValue.value.nightFeeServiceTypesStr = nightFeeServiceTypesArray.value.join(',');
          
          // 如果开启统一价格，将统一价格值填入所有时间段的价格字段
          if (formValue.value.nightFeeUnifiedPrice == 1 && formValue.value.nightFeeUnifiedAmount) {
            nightFeeTimeSlots.value.forEach(timeSlot => {
              timeSlot.price = formValue.value.nightFeeUnifiedAmount;
            });
          }
          formValue.value.nightFeeTimeSlots = JSON.stringify(nightFeeTimeSlots.value);
        } else {
          formValue.value.nightFeeServiceTypesStr = '';
          formValue.value.nightFeeTimeSlots = '';
        }

        // 删除临时数组属性
        const { orderTimeWeekArr, ...submitData } = formValue.value;
        const finalFormValue = { ...submitData };
        updateConfig({ group: group.value, list: finalFormValue }).then((_res) => {
          formBtnLoading.value = false;
          message.success('更新成功');
          load();
        }).catch((_err) => {
          formBtnLoading.value = false;
        });
      } else {
        message.error('验证失败，请填写完整信息');
        formBtnLoading.value = false;
      }
    });
  }

  onMounted(async() => {
    show.value = true;
    await loadCarTypeList()
    await load();
    show.value = false;
  });

  async function load() {
    const res = await getConfig({ group: group.value })
    formValue.value = res.list;
    let orderTimeWeekArr = res.list.orderTimeWeek ? res.list.orderTimeWeek.split(',') : [];
    formValue.value.orderTimeWeekArr = orderTimeWeekArr.map((item) => {
      return parseInt(item)
    })

    orderTimeFormArr.value = res.list.orderTimeForm ? JSON.parse(res.list.orderTimeForm) : [{ startTime: null, endTime: null }]

    orderTimeForm1Arr.value = res.list.orderTimeForm1 ? JSON.parse(res.list.orderTimeForm1) : [{ startTime: null, endTime: null, num: null }]

    formValue.value.maxDateTimeOrderDate = res.list.maxDateTimeOrderDate ? res.list.maxDateTimeOrderDate : null
    orderTimeForm2Arr.value = res.list.orderTimeForm2 ? JSON.parse(res.list.orderTimeForm2) : [{ startTime: null, endTime: null, num: null }]

    // 加载多日期时间段设置数据，如果没有则使用默认示例数据
    if (res.list.maxDateTimeOrderDateTimeJson) {
      dateTimeSettingArr.value = JSON.parse(res.list.maxDateTimeOrderDateTimeJson)
    }

    if (res.list.maxOneDayOrderNumJson) {
      dailyLimits.value = JSON.parse(res.list.maxOneDayOrderNumJson)
    }

    // 加载深夜费时间段配置数据
    if (res.list.nightFeeTimeSlots) {
      nightFeeTimeSlots.value = JSON.parse(res.list.nightFeeTimeSlots)
    } else {
      nightFeeTimeSlots.value = [{ startTime: null, endTime: null, price: null }]
    }

    // 加载深夜费服务类型数据
    if (res.list.nightFeeServiceTypesStr) {
      nightFeeServiceTypesArray.value = res.list.nightFeeServiceTypesStr.split(',');
    } else {
      nightFeeServiceTypesArray.value = [];
    }
  }
</script>
