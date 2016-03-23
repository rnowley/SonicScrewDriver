package org.me.hotel

class GuestTest extends UnitTest("Guest") {

	it should "have its name defined" in {
		Guest("Seline")
		an [IllegalArgumentException] should be thrownBy {
			Guest("")
		}
	}

}
