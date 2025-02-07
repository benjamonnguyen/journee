// Code generated by templ - DO NOT EDIT.

// templ: version: v0.3.819
package journal

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "fmt"

// TODO better design. dispatch status event to status element to maintain LOB
func JournalView() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 1, "<div id=\"app\" x-data=\"{\n            status: {\n                msg: &#39;Synchronized&#39;,\n                cls: &#39;&#39;,\n            },\n            selectedEmotionID: 0,\n            async onDateChange() {\n                this.status.msg = &#39;Loading...&#39;;\n                const resp = await fetch(&#39;/api/journal/&#39;+$refs.date.value);\n\n                if (!resp.ok &amp;&amp; resp.status != 404) {\n                    this.status = { msg: resp.statusText, cls: &#39;status-err&#39; };\n                }\n\n                if (resp.ok) {\n                    const journalData = await resp.json();\n                    $refs.content.value = journalData.content;\n                    this.selectedEmotionID = journalData.emotionID;\n                    setEnergyLevel(journalData.energyLevel);\n                } else {\n                    $refs.content.value = &#39;&#39;;\n                    this.selectedEmotionID = 0;\n                    resetEnergyLevel();\n                }\n\n                this.status.msg = &#39;Synchronized&#39;;\n            },\n            async updateJournal(el) {\n                const formData = new FormData($refs.form);\n                const response = await fetch(`/api/journal/${formData.get(&#39;date&#39;)}/${el}`, {\n                    method: &#39;POST&#39;,\n                    body: formData.get(el),\n                });\n\n                if (response.redirected) {\n                    window.location.href = response.url;\n                    return;\n                }\n\n                if (!response.ok) {\n                    this.status = { msg: (await response.json()).message, cls: &#39;status-err&#39; };\n                } else {\n                    this.status.msg = &#39;Synchronized&#39;;\n                }\n            },\n            showTrackers: true,\n        }\" @keyup.ctrl.e.window.prevent=\"$refs.content.focus()\" @keyup.ctrl.d.window.prevent=\"$refs.date.focus()\" @keyup.ctrl.left.window.prevent=\"console.log(&#39;TODO go next date&#39;)\"><form id=\"journal-form\" x-ref=\"form\"><input id=\"datepicker\" name=\"date\" x-ref=\"date\" @change=\"onDateChange\" x-init=\"$el.value = new Date().toLocaleDateString(&#39;en-CA&#39;); await onDateChange();\"><div id=\"journal-content-container\"><textarea id=\"journal-content\" x-ref=\"content\" name=\"content\" @input=\"status = { msg: &#39;Unsaved changes&#39; }\" @input.throttle.1000ms=\"updateJournal(&#39;content&#39;)\" @input.debounce.1000ms=\"updateJournal(&#39;content&#39;)\" @blur=\"showTrackers = true; $dispatch(&#39;update-slider&#39;)\" @focus=\"\n                    if (isMobile()) {\n                        showTrackers = false;\n                        $refs.date.scrollIntoView({\n                            block: &#39;start&#39;\n                        });\n                    }\"></textarea></div><div id=\"tracker-container\" x-show=\"showTrackers\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = energySlider(1, 100).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = emotionSelection().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 2, "</div><div id=\"form-footer\"><span id=\"status\" x-text=\"status.msg\" :class=\"status.cls\"></span></div></form></div><script>\n        function isMobile() {\n            return window.innerWidth < 768;\n        }\n\n        // datepicker\n        const dateInput = document.querySelector(\"#datepicker\");\n        const dp = new AirDatepicker(\"#datepicker\", {\n            locale: localeEn,\n            onSelect({date, datepicker}) {\n                datepicker.$el.value = date?.toLocaleDateString(\"en-CA\");\n                dateInput.dispatchEvent(new CustomEvent(\"change\"));\n                datepicker.cl\n            },\n            maxDate: dateInput.value,\n            // /** @param {Date} date */\n            // onRenderCell({date, cellType}) {\n            //     // TODO migrate to date type; fetch all in a month\n            //     if (cellType === 'day') {\n            //         return {\n            //             html: '<span class=\"datepicker-day-number\">' + date. + '</span><span class=\"datepicker-dot\"></span>'\n            //         };\n            //     }\n            // },\n            buttons: [\n                {\n                    content: 'Today',\n                    onClick: (dp) => {\n                        let date = new Date();\n                        dp.selectDate(date);\n                        dp.setViewDate(date);\n                    }\n                },\n            ],\n            autoClose: true,\n            isMobile: isMobile(),\n        });\n\n        /**\n        * @typedef {Object} JournalData\n        * @property {string} date\n        * @property {string} content\n        * @property {number} emotionID\n        * @property {number} energyLevel\n        */\n\n        /**\n        * @typedef {Object} Status\n        * @property {string} msg\n        * @property {string} cls\n        */\n    </script><style>\n        #datepicker {\n            margin-bottom: 2px;\n        }\n        #datepicker::-webkit-calendar-picker-indicator{\n            display: none;\n        }\n        #journal-form {\n            width: 100%;\n        }\n        #journal-content-container {\n            margin-bottom: 15px;\n        }\n        @media (max-width: 767px) {\n            #journal-content:focus {\n                height: 50vh;\n            }#journal-content {\n                height: 20vh;\n            }\n        }\n        #journal-content {\n            font-size:12pt;\n            resize: none;\n            width: 100%;\n            height: 40vh;\n            padding: 10px;\n            box-sizing: border-box;\n        }\n        #form-footer {\n            display: flex;\n            align-items: center;\n            flex-direction: column;\n            width: 100%;\n            margin-top: 1.5rem;\n        }\n        #tracker-container {\n            width: 100%;\n            display: flex;\n            align-items: center;\n            flex-direction: column;\n        }\n        #energy-slider-container {\n            margin-bottom: 1.5rem;\n            max-width: 30rem;\n            width: 100%;\n        }\n        #emotion-selection-container {\n            margin-bottom: 1.5rem;\n            max-width: 30rem;\n        }\n        #status {\n            color: var(--grey-800);\n            font-size: small;\n        }\n        #status.status-err {\n            color: var(--red-primary);\n        }\n        .datepicker-dot {\n            display: block;\n            width: 6px;\n            height: 6px;\n            border-radius: 50%;\n            background-color: red; /* You can change the color */\n            margin: 0 auto;\n            margin-top: 5px; /* Adjust the space between the day number and the dot */\n        }\n    </style>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func emotionSelection() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var2 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var2 == nil {
			templ_7745c5c3_Var2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 3, "<style>\n        fieldset {\n            display: flex;\n            padding: 10px;\n            width: fit-content;\n            margin-bottom: 5px;\n        }\n\n        fieldset label {\n            display: inline-block;\n            margin-right: 10px;\n        }\n    </style><div class=\"emotion-selection-container\"><fieldset required><legend style=\"align: center; margin: 0 auto;\">How did you feel overall today?</legend> ")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, opt := range emotionOptions {
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 4, "<div><input type=\"radio\" id=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(opt.getID())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `journal/journal_view.templ`, Line: 241, Col: 22}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 5, "\" value=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var4 string
			templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(opt.getID())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `journal/journal_view.templ`, Line: 242, Col: 25}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 6, "\" name=\"emotion_id\" aria-label=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var5 string
			templ_7745c5c3_Var5, templ_7745c5c3_Err = templ.JoinStringErrs(opt.Name)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `journal/journal_view.templ`, Line: 244, Col: 27}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var5))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 7, "\" x-model=\"selectedEmotionID\" @input=\"status = { msg: &#39;Saving...&#39; }; updateJournal(&#39;emotion_id&#39;)\"> <label for=\"")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var6 string
			templ_7745c5c3_Var6, templ_7745c5c3_Err = templ.JoinStringErrs(opt.getID())
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `journal/journal_view.templ`, Line: 248, Col: 29}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var6))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 8, "\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var7 string
			templ_7745c5c3_Var7, templ_7745c5c3_Err = templ.JoinStringErrs(opt.Emoji)
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `journal/journal_view.templ`, Line: 248, Col: 43}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var7))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 9, "</label></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 10, "</fieldset></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

func energySlider(min, max int) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var8 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var8 == nil {
			templ_7745c5c3_Var8 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 11, "<style>\n        input[type=\"range\"] {\n            width: 100%;\n            height: 6px;\n            -webkit-appearance: none;\n            background: #ddd;\n            border-radius: 4px;\n            outline: none;\n            margin-top: 20px;\n            cursor: pointer;\n        }\n\n        /* Hide default thumb */\n        input[type=\"range\"]::-webkit-slider-thumb {\n            -webkit-appearance: none;\n            width: 30px;\n            height: 30px;\n            background: transparent;\n            border: none;\n            position: relative;\n            outline: none;\n        }\n\n        input[type=\"range\"]::-moz-range-thumb {\n            width: 30px;\n            height: 30px;\n            background: transparent;\n            border: none;\n            position: relative;\n            outline: none;\n        }\n\n        #energy-slider-container {\n            text-align: center;\n        }\n\n        #value-display {\n            font-size: 16px;\n        }\n\n        #slider-thumb {\n            position: absolute;\n            font-size: 22px;\n            transform: translateX(-50%) translateY(-75%);\n            pointer-events: none;\n            user-select: none;\n        }\n    </style><div id=\"energy-slider-container\"><div style=\"position: relative;\"><legend id=\"value-display\">What was your energy level when you woke up?</legend> <input type=\"range\" id=\"energy-slider\" name=\"energy_level\" min=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var9 string
		templ_7745c5c3_Var9, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", min))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `journal/journal_view.templ`, Line: 311, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var9))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 12, "\" max=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var10 string
		templ_7745c5c3_Var10, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprintf("%d", max))
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `journal/journal_view.templ`, Line: 312, Col: 32}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var10))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 13, "\" @input=\"setEnergyLevel($el.value)\" @input.debounce.500ms=\"status = { msg: &#39;Saving...&#39; }; updateJournal(&#39;energy_level&#39;)\" x-init=\"resetEnergyLevel\" @update-slider.window=\"$nextTick(() =&gt; updateSliderThumb())\"><div id=\"slider-thumb\">⚪</div></div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ.JSONScript("energy-thresholds", energyThresholds).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templruntime.WriteString(templ_7745c5c3_Buffer, 14, "<script>\n        const slider = document.getElementById('energy-slider');\n        const sliderThumb = document.getElementById('slider-thumb');\n        const valueDisplay = document.getElementById('value-display');\n\n        const energyThresholds = JSON.parse(document.getElementById('energy-thresholds').textContent);\n\n        function getEmoji(value) {\n            for (let i = 0; i < energyThresholds.length; i++) {\n                if (value <= energyThresholds[i].Threshold) {\n                    return energyThresholds[i].Emoji;\n                }\n            }\n            return energyThresholds[-1].Emoji;\n        }\n\n        function updateSliderThumb() {\n            const percent = slider.value / slider.max;\n            const thumbPosition = percent * (slider.offsetWidth);\n            sliderThumb.style.left = `${thumbPosition}px`;\n        }\n\n        function resetEnergyLevel() {\n            slider.value = 50;\n            valueDisplay.textContent = \"What was your energy level when you woke up?\";\n            sliderThumb.textContent = \"⚪\";\n            updateSliderThumb();\n        }\n\n        function setEnergyLevel(n) {\n            slider.value = n;\n            valueDisplay.textContent = \"Energy level: \" + slider.value + \"%\";\n            sliderThumb.textContent = getEmoji(slider.value);\n            updateSliderThumb();\n        }\n\n        window.addEventListener('resize', updateSliderThumb);\n    </script>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return nil
	})
}

var _ = templruntime.GeneratedTemplate
