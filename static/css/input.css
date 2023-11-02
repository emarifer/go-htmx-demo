@import url('https://fonts.googleapis.com/css2?family=Kanit&display=swap');

@tailwind base;
@tailwind components;
@tailwind utilities;

@layer base {
    body {
        @apply font-Kanit;
    }
}

@keyframes fade-in {
    from {
        opacity: 0;
    }
}

@keyframes fade-out {
    to {
        opacity: 0;
    }
}

@keyframes slide-from-right {
    from {
        transform: translateX(90px);
    }
}

@keyframes slide-to-left {
    to {
        transform: translateX(-90px);
    }
}

/* define animations for the old and new content */
::view-transition-old(slide-it) {
    animation: 180ms cubic-bezier(0.4, 0, 1, 1) both fade-out,
        600ms cubic-bezier(0.4, 0, 0.2, 1) both slide-to-left;
}

::view-transition-new(slide-it) {
    animation: 420ms cubic-bezier(0, 0, 0.2, 1) 90ms both fade-in,
        600ms cubic-bezier(0.4, 0, 0.2, 1) both slide-from-right;
}

/* tie the view transition to a given CSS class */
.sample-transition {
    view-transition-name: slide-it;
}

/***** MODAL DIALOG ****/
#modal {
    /* Underlay covers entire screen. */
    position: fixed;
    top: 0px;
    bottom: 0px;
    left: 0px;
    right: 0px;
    background-color: rgba(0, 0, 0, 0.5);
    z-index: 1000;

    /* Flexbox centers the .modal-content vertically and horizontally */
    display: flex;
    flex-direction: column;
    align-items: center;

    /* Animate when opening */
    animation-name: fadeIn;
    animation-duration: 300ms;
    animation-timing-function: ease;
}

#modal>.modal-underlay {
    /* underlay takes up the entire viewport. This is only
	required if you want to click to dismiss the popup */
    position: absolute;
    z-index: -1;
    top: 0px;
    bottom: 0px;
    left: 0px;
    right: 0px;
}

#modal>.modal-content {
    /* color: black; */
    /* Position visible dialog near the top of the window */
    margin-top: 10vh;

    /* Sizing for visible dialog */
    width: 80%;
    max-width: 400px;

    /* Display properties for visible dialog*/
    /* border: solid 1px #999;
    border-radius: 8px;
    box-shadow: 0px 0px 20px 0px rgba(0, 0, 0, 0.3);
    background-color: white;
    padding: 20px; */

    /* Animate when opening */
    animation-name: zoomIn;
    animation-duration: 300ms;
    animation-timing-function: ease;
}

#modal.closing {
    /* Animate when closing */
    animation-name: fadeOut;
    animation-duration: 300ms;
    animation-timing-function: ease;
}

#modal.closing>.modal-content {
    /* Animate when closing */
    animation-name: zoomOut;
    animation-duration: 300ms;
    animation-timing-function: ease;
}

@keyframes fadeIn {
    0% {
        opacity: 0;
    }

    100% {
        opacity: 1;
    }
}

@keyframes fadeOut {
    0% {
        opacity: 1;
    }

    100% {
        opacity: 0;
    }
}

@keyframes zoomIn {
    0% {
        transform: scale(0.8);
    }

    100% {
        transform: scale(1);
    }
}

@keyframes zoomOut {
    0% {
        transform: scale(1);
    }

    100% {
        transform: scale(0.8);
    }
}