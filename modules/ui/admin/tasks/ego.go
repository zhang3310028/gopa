// Generated by ego.
// DO NOT EDIT

package tasks

import (
	"fmt"
	api "github.com/infinitbyte/gopa/core/http"
	"github.com/infinitbyte/gopa/core/model"
	"github.com/infinitbyte/gopa/modules/ui/admin/common"
	"html"
	"io"
	"net/http"
)

var _ = fmt.Sprint("") // just so that we can keep the fmt import for now
func Index(w io.Writer, r *http.Request, domain string, from, size, taskCount int, tasks []model.Task, domainsCount int, domains []model.Domain) error {
	_, _ = io.WriteString(w, "\n\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n")
	_, _ = io.WriteString(w, "\n\n")
	common.Head(w, "Tasks", "")
	_, _ = io.WriteString(w, "\n<link rel=\"stylesheet\" href=\"/static/assets/css/tasks.css\" />\n<script src=\"/static/assets/js/jquery.timeago.js\"></script>\n<script src=\"/static/assets/js/page/tasks.js\"></script>\n<script src=\"/static/assets/uikit-2.27.1/js/components/pagination.min.js\"></script>\n\n")
	common.Body(w)
	_, _ = io.WriteString(w, "\n")
	common.Nav(w, "Tasks")
	_, _ = io.WriteString(w, "\n\n")

	paras := map[string]interface{}{}
	paras["domain"] = domain

	_, _ = io.WriteString(w, "\n\n<div class=\"tm-middle\">\n\n    <div class=\"uk-container uk-container-center\">\n\n        <div class=\"uk-grid\" data-uk-grid-margin=\"\">\n            <div class=\"tm-sidebar uk-width-medium-1-4 uk-hidden-small uk-row-first\">\n\n                <ul class=\"tm-nav uk-nav\" data-uk-nav=\"\">\n\n                    <!--<li class=\"uk-nav-header\">Tasks</li>-->\n                    <!--<li class=\"uk-active\"><a href=\"#create-task-modal\" data-uk-modal>Create a task</a></li>-->\n                </ul>\n\n            </div>\n            <div class=\"tm-main uk-width-medium-3-4\">\n\n                <article class=\"uk-article\">\n\n\n                </article>\n\n            </div>\n        </div>\n\n        <div class=\"uk-grid\" data-uk-grid-margin>\n\n\n            <div class=\"uk-width-2-10\">\n                <div class=\"uk-alert\" ><span id=\"domain-alert\">Total ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(domainsCount)))
	_, _ = io.WriteString(w, "</span></div>\n                <ul id=\"domain-records\" class=\"uk-list\">\n                    ")
	if len(domains) > 0 {
		for _, domain := range domains {

			_, _ = io.WriteString(w, "\n                    ")
			_, _ = fmt.Fprint(w, GetDomainRow(domain))
			_, _ = io.WriteString(w, "\n                    ")

		}
	}

	_, _ = io.WriteString(w, "\n                </ul>\n            </div>\n\n            <div class=\"uk-width-8-10\">\n                <div class=\"uk-alert\" ><span id=\"alert\">Total ")
	_, _ = io.WriteString(w, html.EscapeString(fmt.Sprint(taskCount)))
	_, _ = io.WriteString(w, "</span></div>\n\n                <div class=\"uk-overflow-container\">\n                    <table id=\"tasks\" class=\"uk-table uk-table-hover uk-table-striped\" cellspacing=\"0\" width=\"100%\">\n                    <thead>\n                    <tr>\n                        <th>URL</th>\n                        <th>LastUpdate</th>\n                        <th>NextCheck</th>\n                        <th>Status</th>\n                    </tr>\n                    </thead>\n                    <tbody id=\"records\">\n                    ")
	if len(tasks) > 0 {
		for _, task := range tasks {

			_, _ = io.WriteString(w, "\n                    ")
			_, _ = fmt.Fprint(w, GetTaskRow(task))
			_, _ = io.WriteString(w, "\n                    ")

		}
	}

	_, _ = io.WriteString(w, "\n                    </tbody>\n                </table>\n\n                    ")
	_, _ = fmt.Fprint(w, api.GetPagination(from, size, taskCount, "", paras))
	_, _ = io.WriteString(w, "\n\n                </div>\n            </div>\n        </div>\n\n    </div>\n\n\n\n</div>\n\n")
	common.OffCanvas(w)
	_, _ = io.WriteString(w, "\n\n<!-- modal start -->\n\n<div id=\"create-task-modal\" class=\"uk-modal\" aria-hidden=\"true\" style=\"display: none; overflow-y: auto;\">\n    <div class=\"uk-modal-dialog uk-modal-dialog-blank\">\n        <button class=\"uk-modal-close uk-close\" type=\"button\"></button>\n        <div class=\"uk-grid uk-flex-middle\" data-uk-grid-margin=\"\">\n            <div class=\"uk-width-medium-1-3 uk-height-viewport uk-cover-background uk-row-first\" style=\"background:#000000;\"></div>\n            <div class=\"uk-width-medium-2-3 uk-row-first \">\n                <h1>Create Task</h1>\n                <div class=\"uk-width-medium-3-4\">\n                    <form class=\"uk-form\">\n                        <fieldset data-uk-margin>\n                            <textarea type=\"text\" placeholder=\"Paste your urls here, one line one url\" rows=\"10\" class=\"uk-width-1-1\"></textarea>\n                            <input type=\"text\" placeholder=\"MaxDepth, eg: 30\">\n                            <input type=\"text\" placeholder=\"MaxBreadth, eg: 3\">\n\n                            <input type=\"checkbox\" id=\"form-s-c\" />\n                            <label for=\"form-s-c\">Follow Sub-domain</label>\n                            <br/>\n                            <input name=\"include-ext\" type=\"checkbox\" id=\"img-ext\" />\n                            <label for=\"img-ext\">Include Images</label>\n                            <input  name=\"include-ext\" type=\"checkbox\" id=\"css-ext\" />\n                            <label for=\"css-ext\">Include Css</label>\n                            <input  name=\"include-ext\" type=\"checkbox\" id=\"js-ext\" />\n                            <label for=\"js-ext\">Include JS</label>\n                        </fieldset>\n                        <button class=\"uk-button\">Create</button>\n                    </form>\n                </div>\n            </div>\n        </div>\n    </div>\n</div>\n\n<!-- modal end -->\n\n")
	common.Footer(w)
	_, _ = io.WriteString(w, "\n")
	return nil
}
