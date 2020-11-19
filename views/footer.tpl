<footer>
    <div class="author">
      Demo Official website:
      <a href="http://{{.Website}}">{{.Website}}</a> /
      Contact me:
      {{if eq .isEmail 1}}
        <a class="email" href="mailto:{{.Email}}">{{.Email}}</a>
      {{else}}
       <a class="email" href="mailto:{{.Email}}">禁止显示</a>
      {{end}}
    </div>
  </footer>