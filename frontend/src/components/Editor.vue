<template>
    <Codemirror class="codemirror" v-model:value="code" :options="cmOptions" border placeholder="test placeholder" :height="660"
        @change="change" />
</template>

<script lang="ts">
import Codemirror from "codemirror-editor-vue3";
import { defineComponent, ref, computed } from 'vue';
import "codemirror/mode/javascript/javascript.js";
import "codemirror/theme/dracula.css";


window.addEventListener('touchstart', function(event) {
    // some logic
    event.preventDefault(); // <-- that should not be used in passive
    // some other magic
});

window.addEventListener('touchmove', function(event) {
    // some logic
    event.preventDefault(); // <-- that should not be used in passive
    // some other magic
});

export default defineComponent({
    props: {
        Data: {
            type: String
        }
    },
    name: "MyCodeMirror",
    components: { Codemirror },
    setup(props) {
        const code = computed({
            get() {
                return props.Data
            },
            set(val) {

            }
        })

        const change = () => {
        }

        return {
            code,
            change,
            cmOptions: {
                mode: "text/javascript", // Language mode
                theme: "dracula", // Theme
                lineNumbers: true, // Show line number
                smartIndent: true, // Smart indent
                indentUnit: 4, // The smart indent unit is 2 spaces in length
                foldGutter: true, // Code folding
                styleActiveLine: true, // Display the style of the selected row
            },
        };
    },
});
</script>

<style>


.codemirror{
    touch-action: none;
}

</style>