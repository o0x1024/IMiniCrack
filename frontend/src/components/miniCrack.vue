<template>
  <div style="padding: 10px;">
    <a-row :gutter="10">
      <a-col>
        <p style="padding: 5px;">小程序包路径：</p>
      </a-col>
      <a-col>
        <a-input id="drop-section" v-model:value="wxPath" style="width: 530px;"
          placeholder="Win小程序包一般在[C:\Users\{username}}\Documents\WeChat Files\Applet\]目录" @dragover.prevent
          @dragenter.prevent @drop="handleDrop" />
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
        <a-input v-model:value="wxId" style="width: 320px;" placeholder="路径中包含会自动识别，扫描大目录时不用填写" />
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
        <a-col :span="4">
          <a-button type="dashed" @click="btnSensitiveScan">敏感信息扫描</a-button>
        </a-col>
      </a-row>



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
                  <a-button @click="btnScan" type="primary" :loading="startscaning">开始扫描</a-button>
                </a-col>
                <a-col>
                  <a-button @click="btnStopScan" :loading="stopscaning" type="primary">停止扫描</a-button>
                </a-col>
                <a-col>
                  <a-button @click="btnSaveRuesult" type="primary">导出结果</a-button>
                </a-col>
                <a-col>
                  <a-button @click="btnDisCodeMirror" type="primary">文件浏览</a-button>
                </a-col>
              </a-row>

              <a-row :gutter="10" style="margin-top: 50px;">
                <a-col style="padding:5px">
                  <span>线程数：</span>
                </a-col>
                <a-col :span="15">
                  <a-slider v-model:value="sliderValue" :min="1" :max="50" />
                </a-col>
                <a-col style="padding: 5px;">
                  <span>{{ sliderValue }}</span>
                </a-col>
              </a-row>
            </a-col>

            <a-col :span="12">
              <a-row :gutter="5">
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
                <a-col :span="4">
                  <a-input v-model:value="desc" placeholder="描述" />
                </a-col>
                <a-col>
                  <a-button @click="onEnableAll" type="primary">Y</a-button>
                </a-col>
                <a-col>
                  <a-button @click="onDisableAll" type="primary">N</a-button>
                </a-col>

              </a-row>
              <a-row style="margin-top: 5px;">
                <a-col>
                  <!--  -->
                  <a-table :columns="columns" :data-source="regexdata.list" size="small" :pagination="false"
                    :scroll="{ x: 220, y: 150 }">
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

                      <template v-if="column.dataIndex === 'Status'">
                        <a-switch @click="switchStatus(record.Id)" v-model:checked="record.Status" />
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

    <a-row v-if="SensitiveScanFlag" style="margin-top: 10px;" :gutter="10">
      <a-col style="padding: 5px;">
        <span>筛选:</span>
      </a-col>
      <a-col>
        <a-select v-model:value="SelectValue" show-search placeholder="选择筛选条件" style="width: 150px"
          :options="selectOptions.list" :filter-option="filterOption" @change="handleChange"></a-select>
      </a-col>
      <!-- <a-col style="padding: 5px;">
        <a-button size="small" @click="btnDelConditionn()" type="primary">删除条件
        </a-button>
      </a-col> -->
    </a-row>

    <a-row style="margin-top: 10px;">
      <a-col :span="24">
        <div>

          <a-table :columns="result_columns" :data-source="sensitiveResult.list" size="small" :pagination="false"
            :scroll="{ x: 220, y: 330 }">
            <template #bodyCell="{ column, record }">
              <template v-if="column.dataIndex === 'operation'">
                <a-row :gutter="10" justify="center">
                  <a-col>
                    <a-button size="small" @click="ToCodeMirror(record.Path, record.LineNo)" type="primary">查询详情
                    </a-button>
                  </a-col>
                  <a-col>
                  </a-col>
                </a-row>
              </template>

              <template v-if="column.dataIndex === 'MatchStr'">
                <template v-for="strs in record.MatchStr">
                  <span>{{ strs }}</span>
                </template>
              </template>

              <template v-if="column.dataIndex === 'Path'">
                <a-button type="link" @click="ToCodeMirror(record.Path, record.LineNo)">{{ record.Path }}</a-button>
              </template>

            </template>
          </a-table>

        </div>
        <!-- <a-textarea v-model:value="logger" id="logtextarea" :change="logChange()" placeholder="" :rows="logRow" /> -->
      </a-col>
    </a-row>


    <!-- <div style="position: absolute; left: 10px; bottom: 0px"> -->
      <div v-html="logger" style="position: absolute; left: 10px; bottom: 0px"></div>
    <!-- </div> -->


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


    <a-drawer title="文件浏览" style="text-align: left;" :width="codewidth" :visible="codevisible" @close="onCodeClose">

      <a-row :gutter="10">
        <a-col>
          <p style="padding: 5px;">文件路径：</p>
        </a-col>
        <a-col>
          <a-input-search v-model:value="disFilePath" style="width: 500px;" enter-button placeholder="选择或输入需要打开的文件路径"
            @search="onSelectCodeFile">
          </a-input-search>
        </a-col>
        <a-col>
          <a-button @click="bntOpenCodeFile()" type="primary">打开文件</a-button>
        </a-col>
      </a-row>
      <a-row>
        <a-col :span="24">
          <MyCodeMirror :Data="disData" :LineNum="codeMirrorLineNum"></MyCodeMirror>
        </a-col>
      </a-row>
    </a-drawer>
  </div>
</template>


<script lang="ts">

import { computed, onMounted, reactive } from 'vue'

import { OpenWxPackDir, OpenDir, OpenDecDir, OpenScanDir, GetDefaultOutPath, SelectOpenFile, OpenDisFile } from "../../wailsjs/go/main/App"
import { Unpack } from "../../wailsjs/go/crack/Crack"
import { GetRegx, AddRegex, DelRegex, ScanSensitive, SaveResult, SaveRegex, UpdateRegex, ChangeRegexStatus, StopScan, DisableAllRegex, EnableAllRegex } from "../../wailsjs/go/scan/Scan"
import { EventsOn } from "../../wailsjs/runtime/runtime"
import { InboxOutlined, UploadOutlined, SearchOutlined, DeleteOutlined, SaveOutlined, EditOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
import { defineComponent, ref } from 'vue';
import MyCodeMirror from './Editor.vue'

const columns = [
  { title: 'ID', dataIndex: 'Id', key: '1', align: 'center' },
  { title: '描述', dataIndex: 'Desc', key: '2', align: 'center', },
  { title: '正则', dataIndex: 'Record', key: '3', align: 'center', ellipsis: true, },
  { title: '状态', dataIndex: 'Status', key: '4', align: 'center', },
  { title: '操作', key: 'operation', fixed: 'right', align: 'center', },
]

const result_columns = [
  { title: '规则名称', dataIndex: 'Desc', key: 'Desc', align: 'center', width: 100, },
  { title: '行号', dataIndex: 'LineNo', key: 'LineNo', align: 'center', width: 80 },
  { title: '匹配结果', dataIndex: 'MatchStr', key: 'MatchStr', align: 'center', width: 400 },
  { title: '文件位置', dataIndex: 'Path', key: 'Path', align: 'center', ellipsis: true, width: 300 },
  { title: '操作', dataIndex: 'operation', key: 'operation', align: 'center', width: 90 },
]

interface DelayLoading {
  delay: number;
}

interface Regex {
  Id: string
  Record: string  //正则
  Desc: string
  Status: boolean
}

interface option {
  value: string
  label: string
}

interface Sensitive {
  Desc: string
  MatchStr: Array<string>
  LineNo: string
  Path: string
}


interface Response {
  Sensitives: Sensitive[]
  Data: string
  Err: string
  Regexs: Regex[]
  FileList: string[]
  Msg: string
}


export default defineComponent({
  components: {
    MyCodeMirror,
    SaveOutlined,
    DeleteOutlined,
    EditOutlined,
    InboxOutlined,
    SearchOutlined,
    UploadOutlined
  },

  setup() {
    let backLog = ''
    let desc = ref('')
    let sliderValue = ref(10)
    let wxId = ref('')
    let disData = ref('')
    let codewidth = ref('900px')
    let curEditReg = ref('')
    let curEditDesc = ref('')
    let CurEditId = ''
    let disFilePath = ref('')
    let logMinHeigth = ref('610')
    let wxPath = ref('')
    let startscaning = ref(false);
    let stopscaning = ref(false);
    let codevisible = ref(false)
    let regex = ref('')
    let editRegVisible = ref(false)
    let regexdata: { list: Regex[] } = reactive({ list: [] });
    let regexs = ref('')
    let outPath = ref('')
    let scanPath = ref('')
    let SelectValue = ref<string>()
    let logRow = ref(25)
    let logger = ref('')
    let defaultOutPath = ""
    let SensitiveScanFlag = ref(false)
    const selectOptions: { list: option[] } = reactive({ list: [] })
    let sensitiveResult: { list: Sensitive[] } = reactive({ list: [] });
    let orginSensitiveResult: { list: Sensitive[] } = reactive({ list: [] });
    let codeMirrorLineNum: any = ref('');





    const onDisableAll = async () => {
      const res = await DisableAllRegex()
      if (res.Code === 200) {
        regexdata.list.forEach(item => {
          item.Status = false
        })
        message.success("更新成功")
      }
    }

    const onEnableAll = async () => {
      const res = await EnableAllRegex()
      console.log(res)
      if (res.Code === 200) {
        console.log(12)
        regexdata.list.forEach(item => {
          item.Status = true
        })

        message.success("更新成功")
      }

    }

    const handleDrop = (event: any) => {
      console.log(event.dataTransfer.files)
      event.preventDefault();
      const file = event.dataTransfer.files[0];
      if (file) {
        wxPath.value = file.path; // 或者 file.webkitRelativePath，取决于你的需求
      }
    }

    const handleChange = (value: string) => {
      SelectValue.value = value
      sensitiveResult.list = orginSensitiveResult.list.filter(item => item.Desc == value)

    };

    const filterOption = (input: string, option: any) => {
      return option.value.toLowerCase().indexOf(input.toLowerCase()) >= 0;
    };

    const btnDelConditionn = () => {
      logger.value = backLog
      SelectValue.value = undefined
    }



    onMounted(() => {
      EventsOn("log", (data: any) => {
        if (data) {
          logger.value = "<p style=\"text-align:left; \">" + data + "<p>"
        }
      })

      EventsOn("scan", (data: Sensitive) => {
        if (data) {
          orginSensitiveResult.list.push(data)
          sensitiveResult.list.push(data)
        }
      })

      EventsOn("scan_dis", (data: any) => {
        if (data) {
          logger.value += "<p style=\"text-align:left; \">" + data + "<p>"
        }
      })

    })

    GetDefaultOutPath().then((result) => {
      if (result) {
        outPath.value = result
        defaultOutPath = result
        scanPath.value = result
      }
    })

    GetRegx().then((result) => {
      if (result.Err) {
        message.error(result.Err)
        return
      }
      if (result.Regexs) {
        regexdata.list = result.Regexs
      }
    })



    const btnSensitiveScan = () => {
      SensitiveScanFlag.value = !SensitiveScanFlag.value
      if (SensitiveScanFlag.value) {
        logMinHeigth.value = '350'
      } else {
        logMinHeigth.value = '610'
      }
    }

    const btnSelectOutPath = (e: any) => {
      OpenDir().then((result) => {
        outPath.value = result
        scanPath.value = result

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
          if (pathSlice.length <= 0) return

          for (let i in pathSlice) {
            if (pathSlice[i].indexOf("wx") != -1) {
              wxId.value = pathSlice[i]
              outPath.value = defaultOutPath + "\\" + wxId.value
              scanPath.value = defaultOutPath + "\\" + wxId.value
              return
            }
          }
        }
      })
    }

    const btnDecExport = () => {
      logger.value = ''
      Unpack(wxPath.value, wxId.value, outPath.value).then((result) => {
        if (result.Err) {
          message.error(result.Err)
        } else {
          message.success(result.Msg)
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


    //敏感信息扫描
    const btnScan = () => {
      sensitiveResult.list = []
      if (scanPath.value === 'C:\\Users\\test\\Documents') {
        message.warning("请检查路径是否正确")
        return
      } else {
        message.info("开始扫描")
      }

      startscaning.value = true
      selectOptions.list.length = 0
      ScanSensitive(scanPath.value, sliderValue.value).then((result) => {
        if (result.Err) {
          message.error(result.Err)
          return
        }

        let ses: Sensitive[] = result.Sensitives
        if (result.Sensitives) {
          console.log(result)
          let descArray: string[] = []
          ses.forEach(element => {
            descArray.push(element.Desc)
          });


          let val = Array.from(new Set(descArray))
          let options: option[] = []
          val.forEach(element => {
            let opt: option = { "value": "", "label": "" }
            opt.value = element
            opt.label = element
            options.push(opt)
          });
          selectOptions.list = options
        }
        message.success(result.Msg)
        startscaning.value = false
      })

    }

    const btnAddRegex = () => {
      let regindex = regexdata.list.findIndex(item => item.Record == regex.value)
      if (regindex == -1) {
        let line: Regex = {
          Id: (Number(regexdata.list[regexdata.list.length - 1].Id) + 1).toString(),
          Record: regex.value,
          Desc: desc.value,
          Status: true
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
      SaveResult(orginSensitiveResult.list).then((result) => {
        if (result) {
          message.info(result)
        }
      })
    }

    const btnSaveRegex = () => {
      SaveRegex().then((result) => {
        if (result.Err) {
          message.error(result)
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
        if (result.Err) {
          message.error(result.Err)
          return
        }
        if (result.Regexs) {
          regexdata.list = result.Regexs
        }
      })
      message.success("更新成功")
      editRegVisible.value = false
    }


    const switchStatus = (id: string) => {
      ChangeRegexStatus(id).then(res => {
        if (res.Err) {
          message.error(res.Err)
        }
        if (res.Msg) {
          message.success(res.Msg)
        }
      })
    }

    const onCodeClose = () => {
      codevisible.value = false
    }

    const btnDisCodeMirror = (lineno: any) => {
      codeMirrorLineNum.value = lineno
      codevisible.value = true
    }


    const onSelectCodeFile = () => {
      SelectOpenFile().then((res: Response) => {
        if (res.Err) {
          message.error(res.Err)
        } else if (res.Data) {
          disFilePath.value = res.Data
          OpenDisFile(disFilePath.value).then((res: Response) => {
            if (res.Err) {
              message.error(res.Err)
            } else if (res.Data) {
              disData.value = res.Data
            }
          })
        }
      })
    }


    const bntOpenCodeFile = () => {
      OpenDisFile(disFilePath.value).then((res: Response) => {
        if (res.Err) {
          message.error(res.Err)
        } else if (res.Data) {
          disData.value = res.Data
        }
      })
    }

    const ToCodeMirror = (path: any, lineno: any) => {
      codevisible.value = true
      disFilePath.value = path
      codeMirrorLineNum.value = lineno

      OpenDisFile(disFilePath.value).then((res: Response) => {
        if (res.Err) {
          message.error(res.Err)
        } else if (res.Data) {
          disData.value = res.Data
        }
      })


    }

    const btnStopScan = () => {
      startscaning.value = false
      stopscaning.value = true
      StopScan().then((res) => {
        if (res) {
          stopscaning.value = false
        }
      })
    }

    return {
      logger,
      scanPath,
      wxId,
      SensitiveScanFlag,
      regex,
      disFilePath,
      onDisableAll,
      disData,
      sliderValue,
      selectOptions,
      result_columns,
      logMinHeigth,
      SelectValue,
      codevisible,
      outPath,
      columns,
      handleDrop,
      codeMirrorLineNum,
      logRow,
      // regxList,
      regexdata,
      codewidth,
      btnStopScan,
      sensitiveResult,
      curEditDesc,
      desc,
      regexs,
      stopscaning,
      startscaning,
      curEditReg,
      wxPath,
      editRegVisible,
      filterOption,
      ToCodeMirror,
      bntOpenCodeFile,
      onEnableAll,
      onSelectCodeFile,
      btnDisCodeMirror,
      onCodeClose,
      switchStatus,
      handleChange,
      btnAddRegex,
      btnScan,
      btnEditRegex,
      btnSaveRuesult,
      btnOpenWxPackDir,
      bntEditOk,
      btnDelConditionn,
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
