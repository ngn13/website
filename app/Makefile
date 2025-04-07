NF_CSS  = static/css/nerdfonts.css
NF_WOFF = static/assets/nerdfonts.woff2

all: $(NF_CSS)
	npm run build

$(NF_CSS): $(NF_WOFF)
	wget "https://www.nerdfonts.com/assets/css/webfont.css" -O $@
	sed 's/\.\.\/fonts\/Symbols-2048-em Nerd Font Complete\.woff2/\/assets\/nerdfonts\.woff2/g' -i $@

$(NF_WOFF):
	wget "https://www.nerdfonts.com/assets/fonts/Symbols-2048-em%20Nerd%20Font%20Complete.woff2" -O $@

run: $(NF_CSS)
	npm run dev

format:
	npm run format

clean:
	rm -rf build
	rm $(NF_CSS)
	rm $(NF_WOFF)

.PHONY: format run clean
