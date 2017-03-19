package main

import "testing"

func TestMinify_LineBreak(t *testing.T) {

	testS := `foo
foo`

	resultS := string(minify([]byte(testS)))

	if resultS != "foo foo"+string(byte(10)) {
		t.Fail()
	}
}

func TestMinify_Tab(t *testing.T) {

	testS := `	foo		foo			`

	resultS := string(minify([]byte(testS)))

	if resultS != "foofoo"+string(byte(10)) {
		t.Fail()
	}
}

func TestMinify_Space(t *testing.T) {

	testS := ` foo  foo   `

	resultS := string(minify([]byte(testS)))

	if resultS != " foo foo "+string(byte(10)) {
		t.Fail()
	}
}

func TestSize(t *testing.T) {

	targetSize := 1136
	bSource := []byte(benchSource)
	originalSize := len(bSource)

	min := minify(bSource)
	minSize := len(min)

	t.Log(originalSize, "->", minSize)

	if targetSize < minSize {
		t.Log("failed to reach target :", targetSize)
		t.Fail()
	}

}

func BenchmarkMinify(b *testing.B) {

	bSource := []byte(benchSource)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		benchOutput = minify(bSource)
	}
}

var benchOutput = []byte{}

const benchSource = `class MrCloser extends HTMLElement {

	static get observedAttributes() {
		return ['closeable', 'closer'];
	}

	attributeChangedCallback(attrName, oldVal, newVal) {
		if (attrName == 'closeable' && !!newVal && !this.closeable) {
			this.closeableID = newVal;
		}
		if (attrName == 'closer' && !!newVal && !this.closer) {
			this.closerID = newVal;
		}
	}

	get closed() {
		return this.closeable.classList.contains('closed');
	}

	set closed(val) {
		if (val) {
			this.closeable.classList.add('closed');
		} else {
			this.closeable.classList.remove('closed');
		}
	}


	toggle() {

		this.closed = !this.closed;

	}


	constructor() {
		super();
	}

	connectedCallback() {
		this.setupClick();
	}

	setupClick() {

		if (!this.closerID || !this.closeableID) {
			return;
		}
		this.closer = document.getElementById(this.closerID);
		this.closeable = document.getElementById(this.closeableID);
		if (!this.closer || !this.closeable) {
			return;
		}
		this.closer.addEventListener('click', e => {
			this.toggle();
		});
		this.closer.classList.add('closer');
		if (location.hash.replace('#', '') == this.closerID) {
			this.closed = false;
		} else {
			this.closed = true;
		}
		
	}
};

window.customElements.define('mr-closer', MrCloser);
`
