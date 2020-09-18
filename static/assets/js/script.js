(function ($) {

	'use strict';

	$(window).on('load', function () {
		$('#preloader').delay(200).fadeOut(350);
		$('body').delay(200).css({ 'overflow': 'visible' });
	});

	/*------------------------------------
		Side Bar
	--------------------------------------*/
	$('.aside-open').on('click', function (event) {
		event.preventDefault();
		$(this).addClass('active');
		$('.sidebar-group').addClass('active');
	});

	$('.sidebar-overlay').on('click', function (event) {
		event.preventDefault();
		$('.sidebar-group').removeClass('active');
		$('.aside-open').removeClass('active');
	});

	$('.close-side').on('click', function (event) {
		event.preventDefault();
		$('.sidebar-group').removeClass('active');
		$('.aside-open').removeClass('active');
	});

	/*------------------------------------
		Search bar
	--------------------------------------*/
	$('.search-open').on('click', function (event) {
		event.preventDefault();
		$('.search-area').addClass('active');
	});

	$('.search-close').on('click', function (event) {
		event.preventDefault();
		$('.search-area').removeClass('active');
	});

	/*------------------------------------
		Mobile Menu
	--------------------------------------*/

	if (jQuery(".mobile-menu-active").length > 0) {
		$('.mobile-menu-active').meanmenu({
			meanMenuContainer: '.mobile-menu-area',
			meanScreenWidth: "1200",
			meanRevealPosition: "right",
		});
	}

	/*-------------------------------------------
	    Sticky Header
	--------------------------------------------- */

	let win = $(window);
	let sticky_id = $(".header-area");
	win.on('scroll', function () {
		let scroll = win.scrollTop();
		if (scroll < 245) {
			sticky_id.removeClass("sticky-header");
		} else {
			sticky_id.addClass("sticky-header");
		}
	});


	/*------------------------------------
        Overlay Close
	--------------------------------------*/

	$(window).scroll(function () {
		if ($(this).scrollTop() !== 0) {
			$('#scrollUp').fadeIn();
		} else {
			$('#scrollUp').fadeOut();
		}
	});

	$('#scrollUp').on('click', function () {
		$("html, body").animate({scrollTop: 0}, 600);
		return false;
	});

	/*------------------------------------
        Spacial Offer Slider
	--------------------------------------*/

	if (jQuery(".spacial-offer-slider").length > 0) {
		$('.spacial-offer-slider').slick({
			slidesToShow: 1,
			slidesToScroll: 1,
			autoplay: true,
			autoplaySpeed: 10000,
			dots: false,
			fade: true,
			arrows: true,
			prevArrow: '<button type="button" class="slick-prev"><i class="fas fa-chevron-left"></i></button>',
			nextArrow: '<button type="button" class="slick-next"><i class="fas fa-chevron-right"></i></button>'
		});
	}

	/*------------------------------------
        Main Home Slider
	--------------------------------------*/

	if (jQuery(".main-home-slider").length > 0) {
		var HomeSlider = $('.main-home-slider');

		HomeSlider.on('init', function (e, slick) {
			var $firstAnimatingElements = $('.home-slider:first-child').find('[data-animation]');
			doAnimations($firstAnimatingElements);
		});

		HomeSlider.on('beforeChange', function (e, slick, currentSlide, nextSlide) {
			var $animatingElements = $('.home-slider[data-slick-index="' + nextSlide + '"]').find('[data-animation]');
			doAnimations($animatingElements);
		});

		HomeSlider.slick({
			slidesToShow: 1,
			slidesToScroll: 1,
			autoplay: true,
			autoplaySpeed: 10000,
			dots: false,
			fade: true,
			arrows: true,
			prevArrow: '<button type="button" class="slick-prev">Prev</button>',
			nextArrow: '<button type="button" class="slick-next">Next</button>'
		});
	}

	/*------------------------------------
        Main Home Slider 3
	--------------------------------------*/

	if (jQuery(".main-home-slider-3").length > 0) {
		var HomeSlider3 = $('.main-home-slider-3');

		HomeSlider3.on('init', function (e, slick) {
			var $firstAnimatingElements = $('.home-slider-3:first-child').find('[data-animation]');
			doAnimations($firstAnimatingElements);
		});

		HomeSlider3.on('beforeChange', function (e, slick, currentSlide, nextSlide) {
			var $animatingElements = $('.home-slider-3[data-slick-index="' + nextSlide + '"]').find('[data-animation]');
			doAnimations($animatingElements);
		});

		HomeSlider3.slick({
			slidesToShow: 3,
			slidesToScroll: 1,
			autoplay: true,
			autoplaySpeed: 10000,
			dots: false,
			arrows: true,
			centerMode: true,
			variableWidth: true,
			prevArrow: '<button type="button" class="slick-prev"><i class="far fa-angle-left"></i></button>',
			nextArrow: '<button type="button" class="slick-next"><i class="far fa-angle-right"></i></button>',
			responsive: [
				{
					breakpoint: 992,
					settings: {
						slidesToShow: 1,
						centerMode: false,
						variableWidth: false,
					}
				},
				{
					breakpoint: 767,
					settings: {
						slidesToShow: 1,
						centerMode: false,
						variableWidth: false,
					}
				}
			]
		});

	}

	/*------------------------------------
        Main Home Slider 4
	--------------------------------------*/

	if (jQuery(".home-slider-4").length > 0) {
		var HomeSlider = $('.home-slider-4');

		HomeSlider.on('init', function (e, slick) {
			var $firstAnimatingElements = $('.home-slide-4:first-child').find('[data-animation]');
			doAnimations($firstAnimatingElements);
		});

		HomeSlider.on('beforeChange', function (e, slick, currentSlide, nextSlide) {
			var $animatingElements = $('.home-slide-4[data-slick-index="' + nextSlide + '"]').find('[data-animation]');
			doAnimations($animatingElements);
		});

		HomeSlider.slick({
			slidesToShow: 1,
			slidesToScroll: 1,
			autoplay: true,
			autoplaySpeed: 10000,
			dots: false,
			fade: true,
			arrows: true,
			prevArrow: '<button type="button" class="slick-prev"><i class="far fa-angle-left"></i></button>',
			nextArrow: '<button type="button" class="slick-next"><i class="far fa-angle-right"></i></button>',
		});
	}

	function doAnimations(elements) {
		var animationEndEvents = 'webkitAnimationEnd mozAnimationEnd MSAnimationEnd oanimationend animationend';
		elements.each(function () {
			var $this = $(this);
			var $animationDelay = $this.data('delay');
			var $animationType = 'animated ' + $this.data('animation');
			$this.css({
				'animation-delay': $animationDelay,
				'-webkit-animation-delay': $animationDelay
			});
			$this.addClass($animationType).one(animationEndEvents, function () {
				$this.removeClass($animationType);
			});
		});
	}

	/*------------------------------------
        Testimonial Slider
	--------------------------------------*/

	if (jQuery(".testimonial-slider").length > 0) {
		$('.testimonial-slider').slick({
			slidesToShow: 2,
			slidesToScroll: 1,
			autoplay: true,
			autoplaySpeed: 10000,
			dots: true,
			arrows: true,
			prevArrow: '<button type="button" class="slick-prev"><i class="fas fa-chevron-left"></i></button>',
			nextArrow: '<button type="button" class="slick-next"><i class="fas fa-chevron-right"></i></button>',
			responsive: [
				{
					breakpoint: 800,
					settings: {
						slidesToShow: 1,
					}
				}
			]
		});
	}


	/*------------------------------------
        Contact Map
	--------------------------------------*/
	if (jQuery(".nice-select").length > 0){
		$('.nice-select').niceSelect();
	}

	/*------------------------------------
        Contact Map
	--------------------------------------*/
	function basicmap() {
		// Basic options for a simple Google Map
		// For more options see: https://developers.google.com/maps/documentation/javascript/reference#MapOptions
		var mapOptions = {
			// How zoomed in you want the map to start at (always required)
			zoom: 11,
			scrollwheel: false,
			// The latitude and longitude to center the map (always required)
			center: new google.maps.LatLng(40.6700, -73.9400), // New York
			// This is where you would paste any style found on Snazzy Maps.
			styles: [{
				"stylers": [{
					"hue": "#AADAFF"
				}]
			}, {
				"featureType": "road",
				"elementType": "labels",
				"stylers": [{
					"visibility": "off"
				}]
			}, {
				"featureType": "road",
				"elementType": "geometry",
				"stylers": [{
					"lightness": 100
				}, {
					"visibility": "simplified"
				}]
			}]
		};
		// Get the HTML DOM element that will contain your map
		// We are using a div with id="map" seen below in the <body>
		var mapElement = document.getElementById('contact-map');

		// Create the Google Map using our element and options defined above
		var map = new google.maps.Map(mapElement, mapOptions);

		// Let's also add a marker while we're at it
		var marker = new google.maps.Marker({
			position: new google.maps.LatLng(40.6700, -73.9400),
			map: map,
			icon: "assets/img/icon/map-marker.png",
			title: 'Biver'
		});
	}
	if ($('#contact-map').length != 0) {
		google.maps.event.addDomListener(window, 'load', basicmap);
	}


})(jQuery);
