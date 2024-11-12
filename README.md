# Can `document.domain` still be used to bypass SOP within same-site cross-origin requests?

`document.domain` allows to overwrite document's domain by any of its superdomains

> Hypothesis: This behavior may lead to SOP bypass (`3rd-party.victim.ru` accesses sensetive data from `victim.ru` within non-preflighted requests)

> Fact: No, it may not. At least in most browsers. Overwritten domain affects nothing from now on. Google Chrome, Mozilla Firefox, and Opera were checked

<body>
    <style>
        body, html {
            margin: 0;
            height: 100%;
            display: flex;
            justify-content: space-between;
        }
        .container {
            display: flex;
            width: 100%;
        }
        .container img {
            max-width: 100%;
            height: auto;
        }
        .left {
            flex: 1;
            display: flex;
            justify-content: flex-start;
            align-items: center;
            padding-left: 10px;
        }
        .right {
            flex: 1;
            display: flex;
            justify-content: flex-end;
            align-items: center;
            padding-right: 10px;
        }
    </style>
    <div class="container">
        <div class="left">
            <img src="assets/chrome-try-steal.png">
        </div>
        <div class="right">
            <img src="assets/chrome-try-steal-impersonate.png">
        </div>
    </div>
    <div class="container">
        <div class="left">
            <img src="assets/firefox-try-steal.png">
        </div>
        <div class="right">
            <img src="assets/firefox-try-steal-impersonate.png">
        </div>
    </div>
</body>
</html>

# Sources

1. https://developer.mozilla.org/en-US/docs/Web/API/Document/domain
2. https://learn.microsoft.com/en-us/deployedge/edge-learnmore-origin-keyed-agent-cluster