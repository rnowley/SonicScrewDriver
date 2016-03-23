package org.me.hotel

import org.scalatest.{FlatSpec, Matchers}

abstract class UnitTest(component: String) extends FlatSpec with Matchers {

	behaviour of component

}
