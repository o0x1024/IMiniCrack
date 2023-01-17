
<template>
  <div style="padding: 10px;">
    <a-row :gutter="10">
      <a-col>
        <p style="padding: 5px;">小程序包路径：</p>
      </a-col>
      <a-col>
        <a-input v-model:value="wxPath" style="width: 530px;"
          placeholder="Win小程序包一般在[C:\Users\{username}}\Documents\WeChat Files\Applet\]目录" />
      </a-col>
      <a-col :span="2">
        <a-button @click="btnOpenWxPackDir">
          <template #icon>
            <search-outlined />
          </template>
          选择目录
        </a-button>
      </a-col>
    </a-row>

    <a-row :gutter="10">
      <a-col>
        <p style="padding: 5px;">WXID：&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p>
      </a-col>
      <a-col>
        <a-input v-model:value="wxId" style="width: 310px;" placeholder="路径中包含会自动识别，未识别到请手动填写" />
      </a-col>
    </a-row>


    <a-row :gutter="10">
      <a-col>
        <p style="padding: 5px;">导出路径：&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</p>
      </a-col>
      <a-col style="padding: 5px;">
        <p class="outPath">{{ outPath }}</p>
      </a-col>
      <a-col>
        <a-button @click="btnSelectOutPath">
          <template #icon>
            <UploadOutlined />
          </template>
          选择目录
        </a-button>
      </a-col>
      <a-col>
        <a-button type="primary" @click="btnOpenDecDir">打开解密目录</a-button>
      </a-col>
    </a-row>



    <a-row justify="space-between">
      <a-row :gutter="10">
        <a-col>
          <a-button type="primary" @click="btnDecExport">解密导出</a-button>
        </a-col>
        <a-col>
          <a-button @click="btnEmpty">清空显示结果</a-button>

        </a-col>
      </a-row>


      <a-col :span="4">
        <a-button type="dashed" @click="btnSensitiveScan">敏感信息扫描</a-button>
      </a-col>
    </a-row>

    <a-row v-if="SensitiveScanFlag" style="margin-top: 10px;">

      <a-col :span="24">
        <div class="sensitive" style="padding: 20px;">

          <a-row justify="space-between">
            <a-col>
              <a-row :gutter="5">
                <a-col>
                  <p style="padding: 5px;">扫描路径：</p>
                </a-col>
                <a-col>
                  <a-input v-model:value="scanPath" style="width: 250px;" placeholder="" />
                </a-col>
                <a-col :span="2">
                  <a-button @click="btnOpenScanDir">
                    <template #icon>
                      <UploadOutlined />
                    </template>
                    选择目录
                  </a-button>
                </a-col>
              </a-row>
              <a-row :gutter="10">
                <a-col>
                  <a-button @click="btnScan" type="primary">开始扫描</a-button>
                </a-col>
                <a-col>
                  <a-button @click="btnSaveRuesult" type="primary">导出结果</a-button>
                </a-col>
              </a-row>
            </a-col>

            <a-col :span="11">
              <a-row :gutter="10">
                <a-col>
                  <a-button @click="btnAddRegex" type="primary">添加正则</a-button>
                </a-col>
                <a-col>
                  <a-tooltip>
                    <template #title>保存正则到本地文件</template>

                    <a-button @click="btnSaveRegex()" type="primary"><template #icon>
                        <save-outlined />
                      </template>
                    </a-button>
                  </a-tooltip>
                </a-col>
                <a-col :span="10">
                  <a-input v-model:value="regex" placeholder="正则" />
                </a-col>
                <a-col :span="6">
                  <a-input v-model:value="desc" placeholder="描述" />
                </a-col>
              </a-row>
              <a-row style="margin-top: 5px;">
                <a-col>
                  <!--  -->
                  <a-table :columns="columns" :data-source="regexdata.list" size="small" :pagination="false"
                    :scroll="{ x: 220, y: 90 }">
                    <template #bodyCell="{ column, record }">
                      <template v-if="column.key === 'operation'">
                        <a-row :gutter="10" justify="center">
                          <a-col>
                            <a-tooltip>
                              <template #title>编辑</template>
                              <a-button size="small" @click="btnEditRegex(record.Id)" type="primary"><template #icon>
                                  <edit-outlined />
                                </template>
                              </a-button>
                            </a-tooltip>
                          </a-col>
                          <a-col>

                            <a-tooltip>
                              <template #title>删除</template>
                              <a-button size="small" @click="btnDelRegex(record.Id)" type="primary"><template #icon>
                                  <delete-outlined />
                                </template>
                              </a-button>
                            </a-tooltip>
                          </a-col>
                        </a-row>
                      </template>
                    </template>
                  </a-table>
                  <!-- <a-textarea v-model:value="regexs" placeholder="regx set" allow-clear :rows="4" /> -->
                </a-col>
              </a-row>
            </a-col>
          </a-row>
        </div>
      </a-col>
    </a-row>

    <a-row v-if="SensitiveScanFlag" style="margin-top: 10px;">
      <a-col style="padding: 5px;">
        <span>筛选:</span>
      </a-col>
      <a-col>
      <a-select v-model:value="SelectValue" show-search placeholder="Select a person" style="width: 150px" :options="options"
        :filter-option="filterOption" @focus="handleFocus" @blur="handleBlur" @change="handleChange"></a-select>
      </a-col>
    </a-row>

    <a-row style="margin-top: 10px;">
      <a-col :span="24">
        <!-- :disabled="true" -->
        <div
          :style="{ 'border-style': 'solid', 'border-width': '1px', 'border-color': 'gray', 'padding': '10px', 'overflow-y': 'scroll', 'width': '100%', 'height': logHeigth + 'px' }">
          <div v-html="logger"></div>
        </div>
        <!-- <a-textarea v-model:value="logger" id="logtextarea" :change="logChange()" placeholder="" :rows="logRow" /> -->
      </a-col>
    </a-row>


    <a-modal v-model:visible="editRegVisible" title="编辑" @ok="bntEditOk" width="900px">
      <a-row>
        <a-col style="padding: 5px;">
          <span>描述：</span>
        </a-col>
        <a-col>
          <a-col :span="20">
            <a-input v-model:value="curEditDesc" placeholder="描述" />
          </a-col>
        </a-col>
      </a-row>

      <a-row style="margin-top: 10px;">
        <a-col style="padding: 5px;">
          <span>正则：</span>
        </a-col>
        <a-col :span="20">
          <a-textarea v-model:value="curEditReg" placeholder="正则" auto-size />
        </a-col>

      </a-row>
    </a-modal>
  </div>


</template>


<script lang="ts">
import { computed, onMounted, reactive } from 'vue'
import { OpenWxPackDir, OpenDir, OpenDecDir, OpenScanDir, GetDefaultOutPath } from "../../wailsjs/go/main/App"
import { Unpack } from "../../wailsjs/go/crack/Crack"
import { GetRegx, AddRegex, DelRegex, ScanSensitive, SaveResult, SaveRegex, UpdateRegex } from "../../wailsjs/go/scan/Scan"
import { EventsOn } from "../../wailsjs/runtime/runtime"
import { InboxOutlined, UploadOutlined, SearchOutlined, DeleteOutlined, SaveOutlined, EditOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
import { defineComponent, ref } from 'vue';
import type { SelectProps } from 'ant-design-vue';


const columns = [
  { title: 'ID', dataIndex: 'Id', key: '1', align: 'center', },
  { title: '描述', dataIndex: 'Desc', key: '2', align: 'center', },
  { title: '正则', dataIndex: 'Record', key: '3', align: 'center', ellipsis: true, },
  { title: '操作', key: 'operation', fixed: 'right', align: 'center', },
]

interface Regx {
  Id: string
  Record: string  //正则
  Desc: string
}


export default defineComponent({
  components: {
    SaveOutlined,
    DeleteOutlined,
    EditOutlined,
    InboxOutlined,
    SearchOutlined,
    UploadOutlined
  },
  setup() {
    let desc = ref('')
    let wxId = ref('')
    let curEditReg = ref('')
    let curEditDesc = ref('')
    let CurEditId = ''
    let logHeigth = ref('550')
    let wxPath = ref('')
    let regex = ref('')
    let editRegVisible = ref(false)
    let regexdata: { list: Regx[] } = reactive({ list: [] });
    let regexs = ref('')
    let outPath = ref('')
    let scanPath = ref('c:\\')
    let SelectValue = ('')
    let logRow = ref(25)
    let logger = ref('')
    let defaultOutPath = ""
    let SensitiveScanFlag = ref(false)

    const options = ref<SelectProps['options']>([
      { value: 'jack', label: 'Jack' },
      { value: 'lucy', label: 'Lucy' },
      { value: 'tom', label: 'Tom' },
    ]);

    const handleChange = (value: string) => {
      console.log(`selected ${value}`);
    };
    const handleBlur = () => {
      console.log('blur');
    };
    const handleFocus = () => {
      console.log('focus');
    };
    const filterOption = (input: string, option: any) => {
      return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
    };


    onMounted(() => {
      EventsOn("log", (data: any) => {
        if (data) {
          logger.value += "<p style=\"text-align:left; \">" + data + "<p>"
        }
      })

      EventsOn("scan", (data: any) => {
        if (data) {
          logger.value += "<p style=\"text-align:left; \">" + data + "<p>"
        }
      })


      GetDefaultOutPath().then((result) => {
        if (result) {
          outPath.value = result
          defaultOutPath = result
        }
      })

      GetRegx().then((result) => {
        if (result) {
          regexdata.list = result
        } else {
          message.warning("获取正则列表失败")
        }
      })

    })

    const btnSensitiveScan = () => {
      SensitiveScanFlag.value = !SensitiveScanFlag.value
      if (SensitiveScanFlag.value) {
        logHeigth.value = '330'
      } else {
        logHeigth.value = '550'
      }
    }

    const btnSelectOutPath = (e: any) => {
      OpenDir().then((result) => {
        outPath.value = result
      })
    }

    const btnOpenWxPackDir = () => {
      OpenWxPackDir(wxPath.value).then((result) => {
        if (result) {
          wxPath.value = result
          let pathSlice = wxPath.value.split("\\")
          if (pathSlice.length <= 0) {
            pathSlice = wxPath.value.split("/")
          }
          console.log(wxPath.value)
          console.log(pathSlice)
          if (pathSlice.length <= 0) return

          for (let i in pathSlice) {
            if (pathSlice[i].indexOf("wx") != -1) {
              wxId.value = pathSlice[i]
              scanPath.value = outPath.value + "\\" + wxId.value
              outPath.value = defaultOutPath + "\\" + wxId.value
              return
            }
          }
        }
      })
    }

    const btnDecExport = () => {
      logger.value = ''
      Unpack(wxPath.value, wxId.value, outPath.value).then((result) => {
        if (result) {
          message.info(result)
        }
      })
    }

    const btnOpenDecDir = () => {
      // let dir = outPath.value + wxId.value
      OpenDecDir(outPath.value).then((result) => {
        if (result) {
          message.info(result)
        }
      })
    }

    const logChange = () => {

      let logtextarea = document.getElementById('logtextarea');

      if (logtextarea == null) {
        return
      }
      logtextarea.scrollTop = logtextarea.scrollHeight;
    }

    const btnEmpty = () => {
      logger.value = ""
    }
    const btnOpenScanDir = () => {
      OpenScanDir(scanPath.value).then((result) => {
        if (result) {
          scanPath.value = result
        }
      })
    }

    const btnScan = () => {
      logger.value = ''
      console.log(scanPath.value)
      ScanSensitive(scanPath.value).then((result) => {
        if (result) {
          message.info(result)
        }
      })
    }

    const btnAddRegex = () => {
      let regindex = regexdata.list.findIndex(item => item.Record == regex.value)
      if (regindex == -1) {
        let line: Regx = {
          Id: (regexdata.list.length + 1).toString(),
          Record: regex.value,
          Desc: desc.value
        }

        regexdata.list.push(line)
        AddRegex(regex.value, desc.value)
        message.success("添加成功")
      } else {
        message.warning("已存在")
      }

    }

    const btnDelRegex = (recordId: string) => {
      regexdata.list = regexdata.list.filter(item => item.Id != recordId)
      DelRegex(recordId)
      message.success("删除成功")
    }

    const btnSaveRuesult = () => {
      SaveResult(logger.value).then((result) => {
        if (result) {
          message.info(result)
        }
      })
    }

    const btnSaveRegex = () => {
      SaveRegex().then((result) => {
        if (result) {
          message.info(result)
        }
      })
    }

    const btnEditRegex = (recordID: string) => {
      editRegVisible.value = true
      CurEditId = recordID
      curEditReg.value = regexdata.list.find(item => item.Id == recordID)?.Record || ''
      curEditDesc.value = regexdata.list.find(item => item.Id == recordID)?.Desc || ''
    }

    const bntEditOk = () => {
      UpdateRegex(CurEditId, curEditDesc.value, curEditReg.value)

      GetRegx().then((result) => {
        if (result) {
          regexdata.list = result
        } else {
          message.warning("获取正则列表失败")
        }
      })
      message.success("更新成功")
      editRegVisible.value = false
    }

    return {
      logger,
      scanPath,
      wxId,
      SensitiveScanFlag,
      regex,
      options,
      logHeigth,
      SelectValue,
      outPath,
      columns,
      logRow,
      // regxList,
      regexdata,
      curEditDesc,
      desc,
      regexs,
      curEditReg,
      wxPath,
      editRegVisible,
      filterOption,
      handleBlur,
      handleFocus,
      handleChange,
      btnAddRegex,
      btnScan,
      btnEditRegex,
      btnSaveRuesult,
      btnOpenWxPackDir,
      bntEditOk,
      btnDelRegex,
      logChange,
      btnOpenScanDir,
      // wxpathChange,
      btnDecExport,
      btnSaveRegex,
      btnOpenDecDir,
      btnSensitiveScan,
      btnEmpty,
      btnSelectOutPath,
    };
  },
});

</script>


<style scoped>
.sensitive {
  border: 1px dashed gray;
}
</style>
